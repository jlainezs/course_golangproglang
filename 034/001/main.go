package main

import "fmt"

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func main() {
	fmt.Printf("1 + 2 = %d\n", mySum(1, 2))
	fmt.Printf("2 + 2 = %d\n", mySum(2, 2))
	fmt.Printf("51 + 7 = %d\n", mySum(51, 7))
}
