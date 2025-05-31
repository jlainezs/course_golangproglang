package main

import "fmt"

// not need to use type, as it is inferred from the value
var packageVar = 1

const packageConst = 2

func main() {
	var mainVar int = 3
	shortDeclaredString := "Hello" // note the := operator
	const mainConst int = 4

	fmt.Println(packageVar, packageConst, mainVar, mainConst, shortDeclaredString)
	packageVar = 5
	fmt.Println("packageVar changed to ", packageVar)
}
