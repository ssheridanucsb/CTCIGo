package arraysandstrings

import "testing"

func TestUniqueString(t *testing.T) {
	test_string := "abcdefg"

	if unique := isUnique(test_string); !unique {
		t.Error("Expected unique string")
	}
}

func TestNonUniqueString(t *testing.T) {
	test_string := "aaabbb"

	if unique := isUnique(test_string); unique {
		t.Error("Expected non unique string")
	}
}

func TestUniqueStringNoDS(t *testing.T) {
	test_string := "abcdefg"

	if unique := isUniqueNoDS(test_string); !unique {
		t.Error("Expected unique string")
	}
}

func TestNonUniqueStringNoDS(t *testing.T) {
	test_string := "aaabbb"

	if unique := isUniqueNoDS(test_string); unique {
		t.Error("Expected non unique string")
	}
}
