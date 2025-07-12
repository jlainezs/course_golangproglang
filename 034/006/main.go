package main

import (
	"examples/mystr"
	"fmt"
	"strings"
)

const s = "Des del moment en què he agafat el seu llibre" +
	"he caigut a terra rodolant de riure. Algun dia espero llegir-lo" +
	"--" +
	"Groucho Marx"

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Printf("%s\t\n", v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n", mystr.Join(xs))
}
