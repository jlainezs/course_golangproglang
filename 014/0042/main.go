package main

import "fmt"

func main() {
	a := [5]int{}

	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	for i, v := range a {
		fmt.Printf("Type %T\t index %v\t value %v\n", v, i, v)
	}
}
