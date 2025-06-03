package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const MAXVALUE = 10
	x, y := rand.Intn(MAXVALUE), rand.Intn(MAXVALUE)

	fmt.Printf("x is %d and y is %d\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is between 4 and 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("None of the previous cases were met")
	}
}
