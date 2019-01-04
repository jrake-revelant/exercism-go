package isogram

import "strings"

// IsIsogram will analyze a word if it only contains each letter once, making it an isogram. It will return true if the word is an isogram.
func IsIsogram(s string) bool {
	// normalize to lower case, not optimized for performance
	s = strings.ToUpper(s)

	//Iterating each rune of all other runes of the string while ignoring the comparison with itself. Also ignoring spaces and hivens
	for _, r1 := range s {
		var counter int
		for _, r2 := range s {

			if (r1 == r2) && !(r1 == '-') && !(r1 == ' ') {
				counter++
			}

		}
		if counter > 1 {
			return false
		}
	}
	return true
}
