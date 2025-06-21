package main

import "fmt"

func foo(x ...int) int {
	r := 0
	for _, v := range x {
		r += v
	}

	return r
}

func bar(x []int) int {
	r := 0
	for _, v := range x {
		r += v
	}

	return r

}

func main() {
	r1 := foo(1, 2, 3, 4)
	a := []int{1, 2, 3, 4}
	r2 := foo(a...)
	r3 := bar(a)

	fmt.Println(r1, r2, r3)
}
