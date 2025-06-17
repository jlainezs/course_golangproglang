package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  uint
	color  string
}

func main() {
	v1 := vehicle{
		engine: engine{false},
		make:   "Seat",
		model:  "600",
		doors:  3,
		color:  "White",
	}
	v2 := vehicle{
		engine: engine{true},
		make:   "Audi",
		model:  "E6",
		doors:  5,
		color:  "Dark gray",
	}

	fmt.Println(v1, v2)

	fmt.Println(v1.model)
	fmt.Println(v2.model)
}
