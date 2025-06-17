package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "John",
		friends:   map[string]int{"Larry": 48, "Jane": 49, "Cheetah": 20},
		favDrinks: []string{"Beer", "Water"},
	}

	fmt.Println(p1)

	for i, v := range p1.friends {
		fmt.Printf("%s\t%v\t", i, v)
	}
	fmt.Println()

	for _, favD := range p1.favDrinks {
		fmt.Printf("%s\t", favD)
	}
	fmt.Println()
}
