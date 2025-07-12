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

func TestUseCount(t *testing.T) {
	type testData struct {
		phrase          string
		expectedLength  int
		expectedContent map[string]int
	}

	data := []testData{
		{"hello world", 2, map[string]int{"hello": 1, "world": 1}},
		{"my little pony", 3, map[string]int{"my": 1, "little": 1, "pony": 1}},
		{"around the world around the world", 3, map[string]int{"around": 2, "the": 2, "world": 2}},
		{"around the world around the", 3, map[string]int{"around": 2, "the": 2, "world": 1}},
	}

	for _, d := range data {
		r := UseCount(d.phrase)
		if len(r) != d.expectedLength {
			t.Errorf("got %d, expected %d", len(r), d.expectedLength)
		}
		for k, v := range d.expectedContent {
			if v != d.expectedContent[k] {
				t.Errorf("got %d, expected %d", v, d.expectedContent[k])
			}
		}
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("hello world")
	}
}

func ExampleUseCount() {
	fmt.Printf("%v\n", UseCount("hello world"))
	// Output:
	// map[hello:1 world:1]
}
