package main

import "fmt"

type Person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := Person{
		firstName:               "Pep",
		lastName:                "Lainez",
		favoriteIceCreamFlavors: []string{"Nata", "Pera"},
	}
	p2 := Person{
		firstName:               "Antoni",
		lastName:                "Garcia",
		favoriteIceCreamFlavors: []string{"Maduixa", "Rom amb panses"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for _, v := range p1.favoriteIceCreamFlavors {
		fmt.Println(p1.firstName, p1.lastName, v)
	}
	for _, v := range p2.favoriteIceCreamFlavors {
		fmt.Println(p2.firstName, p2.lastName, v)
	}
}
