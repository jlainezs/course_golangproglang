package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// slicing: [inclusive:exclusive] both 0 indexed
	a1 := a[:5]
	a2 := a[5:]
	a3 := a[2:7]
	a4 := a[1:6]

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
}
