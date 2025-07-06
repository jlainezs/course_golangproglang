package main

import (
	"fmt"
)

func main() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()

	go func(c <-chan int) {
		fmt.Println(<-c)
	}(cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
