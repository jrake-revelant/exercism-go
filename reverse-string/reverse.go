package reverse

import (
	"strings"
	"unicode/utf8"
)

func String(s string) string {
	var sb strings.Builder
	if utf8.ValidString(s) {
		in := []rune(s)
		for i := len(in) - 1; i >= 0; i-- {
			sb.WriteRune(in[i])
		}
		return sb.String()
	}
	return ""
}
