package main

import "testing"

// TestMySum executes a table test for the main.mySum function
func TestMySum(t *testing.T) {

	// build a structure to hold both, data and expected result.
	type test struct {
		data   []int
		answer int
	}

	// define the tests
	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}

	// run the tests
	for _, tst := range tests {
		x := mySum(tst.data...) // use spread operator to accommodate the slice into a variadic parameter
		if x != tst.answer {
			t.Error("Expected", tst.answer, "Got", x)
		}
	}
}
