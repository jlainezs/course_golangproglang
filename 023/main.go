package main

// note
// constraints should be added to the package
// > go get golang.org/x/exp/constraints

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type myNumbers interface {
	constraints.Integer | constraints.Float
}

func add[T myNumbers](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.1, 2.2))
}
