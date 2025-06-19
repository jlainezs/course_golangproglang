package main

import (
	"fmt"
	"log"
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

// wrapper function for log.println
func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM HERE ", s.String())
}

func main() {
	b := book{
		title: "West with the night",
	}

	var c count = 42

	logInfo(b) // call String() method in book
	logInfo(c) // call String() method in count
}
