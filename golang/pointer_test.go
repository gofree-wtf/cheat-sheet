package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type ABC struct {
	test int
}

func (a *ABC) Frame() {
	fmt.Println("------------")
	fmt.Printf("addr of ptr: %p\n", a)

	fmt.Print("addr of ptr: ")
	fmt.Println(unsafe.Pointer(a))

	fmt.Print("addr of double ptr: ")
	fmt.Println(&a)

	fmt.Print("addr of double ptr: ")
	fmt.Println(unsafe.Pointer(&a))

	var b **ABC = &a

	fmt.Print("addr of double ptr: ")
	fmt.Println(b)

	fmt.Print("addr of double ptr: ")
	fmt.Println(unsafe.Pointer(b))
}

func TestPointer(t *testing.T) {
	test := &ABC{1}
	for i := 0; i < 10; i++ {
		test.Frame()
	}
}
