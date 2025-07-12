package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	r := Years(10)
	if r != 70 {
		t.Error("Expected 70, got ", r)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func ExampleYears() {
	fmt.Printf("%d\n", Years(10))
	// Output:
	// 70
}

func TestYearsTwo(t *testing.T) {
	r := YearsTwo(10)
	if r != 70 {
		t.Error("Expected 70, got ", r)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYearsTwo() {
	fmt.Printf("%d\n", YearsTwo(10))
	// Output:
	// 70
}
