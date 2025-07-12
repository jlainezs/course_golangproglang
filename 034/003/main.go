package main

import (
	"examples/acdc"
	"fmt"
)

func main() {
	fmt.Printf("%d\n", acdc.Sum(1, 2, 3))
	fmt.Printf("%d\n", acdc.Sum(1, 2, 3, 3, 5, 6, 7))
}
