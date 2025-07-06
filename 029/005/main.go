package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Printf("%d\t%v\n", v, ok)

	close(c)
	v, ok = <-c
	fmt.Printf("%d\t%v\n", v, ok)
}
