package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("customErr: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "hello world",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
