package arraysandstrings

import "testing"

func TestStringCompression(t *testing.T) {
	input := "aabcccccaaa"
	outputExpected := "a2b1c5a3"

	if outputActual := stringCompression(input); outputActual != outputExpected {
		t.Errorf("Actual did not equal expected| Actual: %s, Expected: %s", outputActual, outputExpected)
	}
}

func TestStringCompressionNoCompress(t *testing.T) {
	input := "abcd"
	outputExpected := "abcd"

	if outputActual := stringCompression(input); outputActual != outputExpected {
		t.Errorf("Actual did not equal expected| Actual: %s, Expected: %s", outputActual, outputExpected)
	}
}
