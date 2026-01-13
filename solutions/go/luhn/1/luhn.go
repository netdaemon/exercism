package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	trimmed := strings.ReplaceAll(id, " ", "")
	double := false
	total := 0

	if len(trimmed) <= 1 {
		return false
	}

	for i := len(trimmed) - 1; i >= 0; i-- {
		value, isInt := strconv.Atoi(string(trimmed[i]))

		if isInt != nil {
			return false
		}

		if double {
			value = value * 2

			if value > 9 {
				value -= 9
			}
		}

		total += value

		double = !double
	}

	return total%10 == 0
}
