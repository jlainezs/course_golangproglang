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

	m := map[string]Person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for i, person := range m {
		fmt.Printf("%s\t", i)
		for _, fav := range person.favoriteIceCreamFlavors {
			fmt.Printf("%s\t", fav)
		}
		fmt.Println()
	}
}
