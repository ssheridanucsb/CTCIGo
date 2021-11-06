package arraysandstrings

import (
	"strings"
)

func isPalindromePermutation(input string) bool {
	if input == "" {
		return true
	}

	//remove spaces and make lowercase
	cleanString := strings.ReplaceAll(input, " ", "")
	cleanString = strings.ToLower(cleanString)

	//check if there are an even or odd number of characters
	isEven := len(cleanString)%2 == 0

	//use a map to track the characters
	charMap := make(map[rune]int)

	for _, c := range cleanString {
		charMap[c]++
	}

	// if its even there must be an even number of every unique character to have a palindrome permutation
	if isEven {
		for _, val := range charMap {
			if val%2 != 0 {
				return false
			}
		}
	} else {
		// if its odd there can only be one odd character count for there to be a palindrome permutaiton
		// this will be the middle character
		seenOdd := false
		for _, val := range charMap {
			if val%2 != 0 {
				if seenOdd {
					// there cannot be two odd numbers so we must return false
					return false
				}
				// we have seen an odd number yet so we'll flip the flag
				seenOdd = true
			}

		}
	}
	return true
}
