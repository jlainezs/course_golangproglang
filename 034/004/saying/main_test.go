package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Pep")
	if s != "Hello Pep" {
		t.Error("got ", s, "expected ", "Hello Pep")
	}
}

func ExampleGreet() {
	fmt.Printf("%s\n", Greet("Pep"))
	// Output:
	// Hello Pep
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Pep")
	}
}
