package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCast(t *testing.T) {
	var msg interface{}
	msg = true

	if msgStr, ok := msg.(string); ok {
		fmt.Printf("msg is string type: %s\n", msgStr)
	} else if msgBool, ok := msg.(bool); ok {
		fmt.Printf("msg is bool type: %s\n", strconv.FormatBool(msgBool))
	}
}
