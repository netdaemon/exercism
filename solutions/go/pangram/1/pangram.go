package pangram

import "strings"

var letters map[rune]int

func IsPangram(input string) bool {
	lowerInput := strings.ToLower(input)

	for char := 'a'; char < 'z'; char++ {
		if !strings.ContainsRune(lowerInput, char) {
			return false
		}
	}

	return true
}
