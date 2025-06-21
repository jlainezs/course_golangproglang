package main

import "fmt"

func foo() int {
	return 101
}

func bar() (string, int) {
	return "bar", 119
}

func main() {
	a := foo()
	what, b := bar()

	fmt.Println(a, b, what)
}
