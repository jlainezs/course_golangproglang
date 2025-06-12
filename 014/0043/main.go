package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range a {
		fmt.Printf("Type %T\t index %v\t value %v\n", v, i, v)
	}
}
