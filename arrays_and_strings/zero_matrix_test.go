package arraysandstrings

import "testing"

func TestZeroMatrix3x3(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	outputExpected := [][]int{
		{1, 0, 3},
		{0, 0, 0},
		{7, 0, 9},
	}
	zeroMatrix(input, 3, 3)
	if !equalMatrix(input, outputExpected) {
		t.Error("Actual did not equal expected")
	}
}

func TestZeroMatrix3x3MultipleZeros(t *testing.T) {
	input := [][]int{
		{0, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	outputExpected := [][]int{
		{0, 0, 0},
		{0, 5, 0},
		{0, 0, 0},
	}
	zeroMatrix(input, 3, 3)
	if !equalMatrix(input, outputExpected) {
		t.Error("Actual did not equal expected")
	}
}

func TestZeroMatrix3x3MultipleZerosSameRow(t *testing.T) {
	input := [][]int{
		{0, 0, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	outputExpected := [][]int{
		{0, 0, 0},
		{0, 0, 6},
		{0, 0, 9},
	}
	zeroMatrix(input, 3, 3)
	if !equalMatrix(input, outputExpected) {
		t.Error("Actual did not equal expected")
	}
}
