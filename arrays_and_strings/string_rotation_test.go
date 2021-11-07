package arraysandstrings

import "testing"

func TestStringRotation(t *testing.T) {

	s1 := "waterbottle"
	s2 := "erbottlewat"

	if !isRotation(s1, s2) {
		t.Error("Expect is rotation to be true, but got false")

	}
}
