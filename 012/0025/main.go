package main

import (
	"fmt"
	"math/rand"
)

func myIntN() int {
	return rand.Intn(250)
}

func main() {
	x := myIntN()
	fmt.Printf("x is %d\n", x)
	switch {
	case x >= 0 && x <= 100:
		fmt.Println("x is between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("x is between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("x is between 201 and 250")
	default:
		fmt.Println("It should not show this message.")
	}
}
