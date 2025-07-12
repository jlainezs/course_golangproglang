package acdc

import "fmt"

// ExampleSum is also a test. It will be included in the project godoc
// as an example of the Sum function.
// Additionally, as we are using a Println, we can document
// with the Output comment, the expected output
func ExampleSum() {
	fmt.Println(Sum(1, 2))
	// Output:
	// 3
}
