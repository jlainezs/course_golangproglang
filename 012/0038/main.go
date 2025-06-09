package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := 0
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			c++
			fmt.Printf("Iteration %v\t Count %v\t x is 3\n", i, c)
		}
	}
}
