package main

import "fmt"

func add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.1, 2.2))
}
