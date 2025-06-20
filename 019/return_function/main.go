package main

import "fmt"

func main() {
	x := foo
	fmt.Println(x())

	y := bar()
	r := y()
	fmt.Println(r)
}

func foo() int {
	return 42
}

// a function which returns a function
func bar() func() int {
	return func() int {
		return 43
	}
}
