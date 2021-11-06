package arraysandstrings

func isUnique(input string) bool {
	//takes input string, checks if all characters are unique

	//check for empty string
	if len(input) == 0 {
		return true
	}

	// use map to keep track of seen characters
	charMap := make(map[rune]bool)

	// design question: what do we consider a char - do spaces/punctuation count?
	// iterate over runes in the string
	for _, c := range input {

		if _, ok := charMap[c]; ok {
			return false
		}

		charMap[c] = true

	}
	return true
}

func isUniqueNoDS(input string) bool {
	//checks if string is unique without using any additional data structures

	for index, char_one := range input {
		if index == len(input)-1 {
			break
		}
		for _, char_two := range input[index+1:] {
			if char_one == char_two {
				return false
			}

		}
	}
	return true
}
