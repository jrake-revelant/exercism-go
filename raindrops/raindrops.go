package raindrops

import (
	"strconv"
	"strings"
)

// Convert expects an integer and will return "raindrops" as a pattern based on if the numbers factor contain 3 and/or 5 and/or 7 (see Readme). Since numbers can have multiple factors strings are joined.
func Convert(a int) string {
	var drop string
	if v := a % 3; v == 0 {
		drop = "Pling"
	}
	if v := a % 5; v == 0 {
		drop = strings.Join([]string{drop, "Plang"}, "")
	}
	if v := a % 7; v == 0 {
		drop = strings.Join([]string{drop, "Plong"}, "")
	}
	if (a%3) != 0 && (a%5) != 0 && (a%7) != 0 {
		drop = strconv.Itoa(a)
	}
	return drop
}
