package main

import "fmt"

func sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(nums...))
}
