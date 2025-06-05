package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		description := ""

		switch x {
		case 0:
			description = "zero"
		case 1:
			description = "one"
		case 2:
			description = "two"
		case 3:
			description = "three"
		default:
			description = "four"
		}

		fmt.Printf("Iteration %v\tx value is %s\n", i, description)
	}
}
