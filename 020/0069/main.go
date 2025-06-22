package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("Inside an anonymous function called when the app is running.")
	}

	a()
}
