package arraysandstrings

import "testing"

func equalMatrix(m1 [][]int, m2 [][]int) bool {
	equal := true

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if m1[i][j] != m2[i][j] {
				equal = false
				break
			}
		}
	}
	return equal

}

func TestRotateMatrix2x2(t *testing.T) {
	input := [][]int{
		{1, 0},
		{0, 1},
	}

	outputExpected := [][]int{
		{0, 1},
		{1, 0},
	}

	rotateMatrix(input)

	if !equalMatrix(input, outputExpected) {
		t.Error("Actual did not equal expected")
	}
}

func TestRotateMatrix3x3(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	outputExpected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	rotateMatrix(input)

	if !equalMatrix(input, outputExpected) {
		t.Error("Actual did not equal expected")
	}
}

func TestRotateMatrix4x4(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	outputExpected := [][]int{
		{13, 9, 5, 1},
		{14, 10, 6, 2},
		{15, 11, 7, 3},
		{16, 12, 8, 4},
	}

	rotateMatrix(input)

	if !equalMatrix(input, outputExpected) {
		t.Error("Actual did not equal expected")
	}
}
