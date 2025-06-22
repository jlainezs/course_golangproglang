package main

import "fmt"

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, num int) string {
	return fmt.Sprintf("%d", f(num))
}

func main() {
	fmt.Println(printSquare(square, 10))
}
