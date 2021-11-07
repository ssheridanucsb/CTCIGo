package arraysandstrings

import (
	"strconv"
	"strings"
)

func stringCompression(input string) string {

	if len(input) == 0 {
		return input
	}

	charArray := []rune(input)

	var sb strings.Builder

	count := 1
	for i := 1; i < len(charArray); i++ {
		if charArray[i] == charArray[i-1] {
			count++
		} else {
			sb.WriteRune(charArray[i-1])
			num := strconv.Itoa(count)
			sb.WriteString(num)
			count = 1
		}
	}

	sb.WriteRune(charArray[len(charArray)-1])
	num := strconv.Itoa(count)
	sb.WriteString(num)

	output := sb.String()
	if len(output) > len(input) {
		return input
	}
	return output
}
