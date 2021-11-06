package arraysandstrings

import "testing"

func TestURLifyOneSpace(t *testing.T) {
	input := "abc efg  "
	expected := "abc%20efg"

	if actual := URLify(input, 7); expected != actual {
		t.Errorf("Expected did not equal actual: %s | expected: %s", actual, expected)

	}
}

func TestURLifyTwoSpace(t *testing.T) {
	input := "Mr John Smith    "
	expected := "Mr%20John%20Smith"

	if actual := URLify(input, 13); expected != actual {
		t.Errorf("Expected did not equal actual: %s | expected: %s", actual, expected)

	}
}
