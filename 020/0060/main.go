package main

/*
	Expected execution: inverse of flow...
	On main
	On b
	On a
*/
import "fmt"

func a() {
	fmt.Println("On a")
}

func b() {
	fmt.Println("On b")
}

func main() {
	defer a()
	defer b()
	defer func() {
		fmt.Println("On main")
	}()
}
