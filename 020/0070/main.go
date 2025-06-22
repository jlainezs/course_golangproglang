package main

import "fmt"

func retFunc() func() int {
	return func() int {
		return 42
	}
}

func main() {
	myFunc := retFunc()
	fmt.Println(myFunc())
}
