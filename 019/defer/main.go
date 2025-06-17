package main

import "fmt"

func end() {
	fmt.Println("This is the last text shown by the app.")
}

func main() {
	defer end()
	fmt.Println("main called.")
}
