package arraysandstrings

import (
	"testing"
)

func TestPalindromePermutationOdd(t *testing.T) {
	input := "Tact Coa"
	tf := isPalindromePermutation(input)
	if !tf {
		t.Errorf("Expected palidrome permutation")
	}
}

func TestPalindromePermutationEven(t *testing.T) {
	input := "ppoo"
	tf := isPalindromePermutation(input)
	if !tf {
		t.Errorf("Expected palidrome permutation")
	}
}
