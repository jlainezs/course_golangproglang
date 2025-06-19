package main

import "fmt"

type person struct {
	name string
}

// use a receiver to add methods to a type
func (p person) speak() {
	fmt.Println("speaking to", p.name)
}

func main() {
	p1 := person{
		name: "Bob",
	}

	// use the method speak defined in the struct person
	p1.speak()
}
