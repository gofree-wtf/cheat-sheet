package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
Test Results

Case 1.
- MBP 16' 2019
- Core i9 2.4Ghz 8C16T
- Request rate: 49441.20 req/s

Case 2.
- M1 Mac Mini 2021
- Apple M1 8C8T
- Rosetta 2
- Request rate: 94150.40 req/s
*/

func TestHttpMultipleRequest(t *testing.T) {
	const (
		requestDuration = 5 * time.Second
		port            = 8080
		goroutineCount  = 100
	)

	severCtx, serverStop := context.WithCancel(context.Background())
	go startServer(severCtx, port)
	defer serverStop()

	var totalSuccessCount, totalFailCount uint32
	clientCtx, clientStop := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			fmt.Printf("client %d: start request\n", i)
			successCount, failCount := startRequest(clientCtx, port)
			fmt.Printf("client %d: end request - success count: %d, fail count: %d\n", i, successCount, failCount)

			atomic.AddUint32(&totalSuccessCount, successCount)
			atomic.AddUint32(&totalFailCount, failCount)
		}(i)
	}

	time.Sleep(requestDuration)
	clientStop()
	wg.Wait()

	fmt.Printf("main: total request success count: %d, fail count: %d\n", totalSuccessCount, totalFailCount)
	fmt.Printf("main: request rate: %.2f req/s\n", float64(totalSuccessCount)/requestDuration.Seconds())
}

func startServer(ctx context.Context, port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			fmt.Println(err)
		}
	})
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		fmt.Println("server: start server")
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	<-ctx.Done()
	_ = server.Close()
}

func startRequest(ctx context.Context, port int) (successCount, failCount uint32) {
	url := fmt.Sprintf("http://localhost:%d", port)
	client := &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: 10000,
		},
	}
	stoppedNotifyCh := make(chan struct{})
	run := true

	go func() {
		for run {
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				failCount++
				continue
			}

			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				failCount++
			} else {
				successCount++
			}
			_ = resp.Body.Close()
		}
		stoppedNotifyCh <- struct{}{}
	}()

	<-ctx.Done()
	run = false
	<-stoppedNotifyCh
	return
}
