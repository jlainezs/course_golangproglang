package mystr

import (
	"strings"
	"testing"
)

const s = "Des del moment en què he agafat el seu llibre" +
	"he caigut a terra rodolant de riure. Algun dia espero llegir-lo" +
	"--" +
	"Groucho Marx"

var sps = strings.Split(s, " ")

func BenchmarkCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cat(sps)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join(sps)
	}
}

// to tun the benchmark, open a terminal in the mystr directory
// go test -bench .
// goos: windows
// goarch: amd64
// pkg: examples/mystr
// cpu: AMD Ryzen 9 5900X 12-Core Processor
// BenchmarkCat-24           853132              1218 ns/op
// BenchmarkJoin-24         8050362               149.4 ns/op
// PASS
// ok      examples/mystr  2.613s
//
// From the result we see Join is more efficient than Cat
