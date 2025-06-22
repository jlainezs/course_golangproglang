package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Inside an anonymous function called when the app closes.")
	}()

	func() {
		fmt.Println("Inside an anonymous function called when the app is running.")
	}()
}
