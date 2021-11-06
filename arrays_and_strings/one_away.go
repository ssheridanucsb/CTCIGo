package arraysandstrings

func OneAway(inputOne string, inputTwo string) bool {
	// checks if the two strings are one character apart

	//cast to rune arrays
	r1 := []rune(inputOne)
	r2 := []rune(inputTwo)

	l1 := len(r1)
	l2 := len(r2)

	//swap so l1 is always greater than or equal to
	if l1 < l2 {
		r1, r2 = r2, r1
		l1, l2 = l2, l1
	}
	//check insert/delete
	if l1 == l2+1 {
		missingChar := false

		long_index := 0
		short_index := 0
		for {
			if short_index == l2-1 {
				if ((r1[long_index] != r2[short_index]) && !missingChar) || r1[long_index] == r2[short_index] {
					return true
				}
				return false
			}

			if r1[long_index] != r2[short_index] {
				if missingChar {
					return false
				}
				missingChar = true
				long_index++
			} else {
				short_index++
				long_index++
			}

		}

	} else if l1 == l2 {
		// check letter change
		diffLetter := false
		for i, c := range r1 {
			if c != r2[i] {
				if diffLetter {
					return false
				}
				diffLetter = true
			}
		}
		return true
	}
	return false
}
