package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	go func(c <-chan int) {
		fmt.Println(<-cs)
	}(cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
