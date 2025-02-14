package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(15, 15)
	if total != 30 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
