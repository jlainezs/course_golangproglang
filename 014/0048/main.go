package main

import "fmt"

func main() {
	s0 := []string{"James", "Bond", "Shaken, not stirred"}
	s1 := []string{"Miss", "Moneypenny", "I'm 008."}
	s := [][]string{s0, s1}

	for _, row := range s {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}
