package arraysandstrings

import "strings"

func isRotation(s1 string, s2 string) bool {
	//given two strings and strings.Contains check if s2 is a rotation of s1
	if len(s1) == len(s2) {
		s1s2 := s1 + s2
		return strings.Contains(s1s2, s2)
	}

	return false
}
