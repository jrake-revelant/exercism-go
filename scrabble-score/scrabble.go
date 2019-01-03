package scrabble

import "strings"

// We define the score mapping via a map of runes
var scrabbleScores = map[rune]int{
	'A': 1,
	'E': 1,
	'I': 1,
	'O': 1,
	'U': 1,
	'L': 1,
	'N': 1,
	'R': 1,
	'S': 1,
	'T': 1,
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

// Score will compute a words (passed as string) scrabble score (returned as int), this works for upper and lowercase but is only test for a subset of runes that have a mapping according to the README
func Score(s string) int {
	var score int
	// Given our map of scores is only for uppercases we normalize the input string to uppercase
	var upperS = strings.ToUpper(s)

	for _, value := range upperS {
		score = score + scrabbleScores[value]
	}

	return score
}
