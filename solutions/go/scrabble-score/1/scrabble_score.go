package scrabble

import "unicode"

func Score(word string) int {
	var score int
	for _, letter := range word {
		switch unicode.ToUpper(letter) {
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		default:
			score += 1
		}
	}

	return score
}
