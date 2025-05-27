package main

import "fmt"

func main() {
	s, i, f := "Hello", 42, 3.14159265359
	fmt.Printf("%v is of type %T\n", s, s)
	fmt.Printf("%v is of type %T\n", i, i)
	fmt.Printf("%v is of type %T\n", f, f)
}
