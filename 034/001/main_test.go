package main

import "testing"

// TestMySum tests the main.mySum
// In go tests, by convention, starts with Test followed by the name
// of the function they are testing.
func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}
