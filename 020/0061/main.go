package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("I am %s and I am %d years old.\n", p.first, p.age)
}

func main() {
	p := person{
		first: "Sunkawakan",
		age:   45,
	}

	p.speak()
}
