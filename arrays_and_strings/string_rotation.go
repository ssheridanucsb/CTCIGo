package arraysandstrings

import "strings"

func isRotation(s1 string, s2 string) bool {
	//given two strings and strings.Contains check if s2 is a rotation of s1
	n := len(s1)

	cut := 0
	for i := 0; i < n; i++ {
		if s2[0] == s1[i] {
			cut = i
			break
		}
	}

	return strings.Contains(s2, s1[cut:])
}
