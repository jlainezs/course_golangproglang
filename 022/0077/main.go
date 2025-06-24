package main

import "fmt"

type person struct {
	first string
}

func changeValue(who person, newValue string) person {
	who.first = newValue
	return who
}

func changePtr(who *person, newValue string) {
	who.first = newValue
}

func main() {
	p := person{
		first: "Perica",
	}
	fmt.Println(p)

	otherP := changeValue(p, "Pollo")
	fmt.Println(p, otherP)

	changePtr(&p, "Pastrami")
	fmt.Println(p)
}
