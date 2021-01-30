package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCast(t *testing.T) {
	var msg interface{}
	msg = "test"

	// Option 1.
	switch msg := msg.(type) {
	case string:
		fmt.Printf("msg is string type: %s\n", msg)
	case bool:
		fmt.Printf("msg is bool type: %s\n", strconv.FormatBool(msg))
	}

	// Option 2.
	if msgStr, ok := msg.(string); ok {
		fmt.Printf("msg is string type: %s\n", msgStr)
	} else if msgBool, ok := msg.(bool); ok {
		fmt.Printf("msg is bool type: %s\n", strconv.FormatBool(msgBool))
	}
}
