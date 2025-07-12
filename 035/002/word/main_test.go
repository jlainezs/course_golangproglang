package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	r := Count("hello world")
	if r != 2 {
		t.Errorf("got %d, expected %d", r, 2)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("hello world")
	}
}

func ExampleCount() {
	fmt.Printf("%d\n", Count("hello world"))
	// Output:
	// 2
}
