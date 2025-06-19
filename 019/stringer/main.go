package main

import (
	"fmt"
	"strconv"
)

type count int

type book struct {
	title string
}

// replace stringer.String() method for count type
func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

// replace stringer.String() method for book type
func (b book) String() string {
	return fmt.Sprint("The book is ", b.title)
}

func main() {
	b := book{
		title: "West with the night",
	}

	var c count = 42

	fmt.Println(b) // call String() method in book
	fmt.Println(c) // call String() method in count
}
