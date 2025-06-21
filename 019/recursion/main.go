package main

/*
	recursive factorial operation
*/
import "fmt"

func main() {
	r := factorial(4)
	fmt.Println(r)
}

func factorial(a int) int {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}
