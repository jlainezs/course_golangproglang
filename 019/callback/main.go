package main

import "fmt"

func main() {
	x := doMath(30, 40, add)
	fmt.Println(x)

	y := doMath(20, 10, substract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(x int, y int) int {
	return x + y
}

func substract(x int, y int) int {
	return x - y
}
