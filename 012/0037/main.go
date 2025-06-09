package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("Index %s \t Value %v \n", k, v)
	}

	age := m["James"]
	fmt.Println(age)

	v, ok := m["Q"]
	if !ok {
		fmt.Printf("%v \t Q is not stored in the map", v)
	}
}
