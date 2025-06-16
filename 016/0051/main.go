package main

import "fmt"

func main() {
	a := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}

	a["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	delete(a, "no_dr")

	for k, v := range a {
		fmt.Printf("%s: \t", k)
		for i, w := range v {
			fmt.Printf("%v - %s\t", i, w)
		}
		fmt.Println()
	}
}
