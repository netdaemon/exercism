package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	var total int

	for i, r := range isbn {
		multiplier := i + 1

		if i == 9 && r == 'X' {
			total += 10 * multiplier
			continue
		}

		digit, err := strconv.Atoi(string(r))

		if err != nil {
			return false
		}

		total += digit * multiplier
	}

	return total%11 == 0
}
