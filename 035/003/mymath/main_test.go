package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type testData struct {
		values []int
		result float64
	}
	data := []testData{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000}, 9},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
	}

	for _, v := range data {
		avg := CenteredAvg(v.values)
		if avg != v.result {
			t.Errorf("got %v, expected %v", avg, v.result)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	a := []int{1, 4, 6, 8, 100}
	for i := 0; i < b.N; i++ {
		CenteredAvg(a)
	}
}

func ExampleCenteredAvg() {
	a := []int{1, 4, 6, 8, 100}
	avg := CenteredAvg(a)
	fmt.Printf("%v\n", avg)
	// Output:
	// 6
}
