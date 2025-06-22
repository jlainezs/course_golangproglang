package main

import "fmt"

func counter(a int) func() int {
	var x int = a
	return func() int {
		x++
		return x
	}
}

func main() {
	x := counter(10)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}
