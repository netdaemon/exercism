package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	lower := strings.ToLower(word)

	for i, rune := range lower {
		if unicode.IsLetter(rune) && strings.ContainsRune(lower[i+1:], rune) {
			return false
		}
	}

	return true
}
