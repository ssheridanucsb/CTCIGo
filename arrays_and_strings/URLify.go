package arraysandstrings

import "unicode"

func URLify(input string, input_len int) string {
	// take in a string and replaces spaces with %20
	// assumes there is enough space at the end for an in place replacement
	// input_len is the number of actual runes
	r := []rune("%20") //cast replacement value to array of runes
	s := []rune(input) //cast input string to array of runes
	lr := len(r)
	ls := len(s)

	// iterate over input rune array
	for i := 0; i < input_len; i++ {
		//check if the character is a space
		if unicode.IsSpace(s[i]) {

			// shift array over
			for j := ls - 1; j > i+1; j-- {
				s[j-lr+1], s[j] = s[j], s[j-lr+1]
			}

			// insert %20
			s[i] = r[0]
			for k := 1; k < lr; k++ {
				s[i+k] = r[k]
			}

		}
	}
	//convert array back to string and return
	return string(s)

}
