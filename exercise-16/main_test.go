package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(5, 5)
	if sum != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 10)
	}
}

func TestSubtract(t *testing.T) {
	sum := Subtract(10, 5)
	if sum != 5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 5)
	}
}
