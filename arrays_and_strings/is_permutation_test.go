package arraysandstrings

import "testing"

func TestPermutation(t *testing.T) {
	inputOne := "abcdefg"
	inputTwo := "gfedcba"

	if perm := isPermutation(inputOne, inputTwo); !perm {
		t.Error("Expected permutation")
	}
}

func TestNotPermutation(t *testing.T) {
	inputOne := "abc"
	inputTwo := "yuoue"

	if perm := isPermutation(inputOne, inputTwo); perm {
		t.Error("Expected not permutation")
	}
}
