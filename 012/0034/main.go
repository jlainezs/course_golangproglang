package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("Outer loop %d \t inner loop %d \n", i, j)
		}

		fmt.Printf("\n")
	}
}
