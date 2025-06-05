package main

import "fmt"

func main() {
	i, limit, breakOn := 0, 10, 7

	for i < limit {
		fmt.Println(i)
		i++
		if i > breakOn {
			break
		}
	}
}
