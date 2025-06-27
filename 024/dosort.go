package main

import (
	"fmt"
	"sort"
)

type ppl struct {
	first string
	age   int
}

type ByName []ppl

func (p ByName) Len() int {
	return len(p)
}
func (p ByName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p ByName) Less(i, j int) bool {
	return p[i].first < p[j].first
}

func main() {
	p1 := ppl{"James", 32}
	p2 := ppl{"Moneypenny", 27}
	p3 := ppl{"Q", 64}
	p4 := ppl{"M", 56}

	people := []ppl{p1, p2, p3, p4}
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}
