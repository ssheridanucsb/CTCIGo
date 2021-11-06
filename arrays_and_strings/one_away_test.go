package arraysandstrings

import "testing"

func TestOneAwayDelete(t *testing.T) {
	inputOne := "pale"
	inputTwo := "ple"
	if !OneAway(inputOne, inputTwo) {
		t.Errorf("Expected strings to be one away")
	}
}

func TestOneAwayInsert(t *testing.T) {
	inputOne := "pale"
	inputTwo := "pales"
	if !OneAway(inputOne, inputTwo) {
		t.Errorf("Expected strings to be one away")
	}
}

func TestOneAwayDiff(t *testing.T) {
	inputOne := "pale"
	inputTwo := "bake"
	if OneAway(inputOne, inputTwo) {
		t.Errorf("Expected strings to not be one away")
	}
}
