package arraysandstrings

func isPermutation(inputOne string, inputTwo string) bool {
	// given two input strings check if they're permutations of eachother

	// validate that the strings exists
	if len(inputOne) == 0 && len(inputTwo) == 0 {
		return true
	}

	if len(inputOne) == 0 || len(inputTwo) == 0 {
		return false
	}

	//convert input strings to mapss

	mapOne := make(map[rune]int)
	for _, c := range inputOne {
		mapOne[c] += 1
	}

	mapTwo := make(map[rune]int)
	for _, c := range inputTwo {
		mapTwo[c] += 1
	}

	for _, c := range inputOne {

		countOne := mapOne[c]
		countTwo, ok := mapTwo[c]

		if !ok {
			return false
		}

		if countOne != countTwo {
			return false
		}

	}
	return true
}
