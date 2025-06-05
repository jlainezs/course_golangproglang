package main

import (
	"fmt"
	"math/rand"
)

func main0028() {
	const MAXVALUE = 10
	x, y := rand.Intn(MAXVALUE), rand.Intn(MAXVALUE)

	fmt.Printf("x is %d and y is %d\n", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is between 4 and 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("None of the previous cases were met")
	}
}

func main() {
	for i := 0; i <= 99; i++ {
		fmt.Printf("%d \n", i)
	}

	for i := 0; i <= 99; i++ {
		main0028()
	}
}
