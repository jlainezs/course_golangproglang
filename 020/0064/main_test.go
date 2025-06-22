package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(1, 1)
	want := 2
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(1, 1)
	want := 0

	if got != want {
		t.Errorf("Substract was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestDoMath(t *testing.T) {
	got := doMath(1, 2, add)
	want := 3
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}

	got = doMath(2, 1, subtract)
	want = 1
	if got != want {
		t.Errorf("Substract was incorrect, got: %d, want: %d.", got, want)
	}
}
