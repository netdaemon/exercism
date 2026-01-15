package romannumerals

import (
	"errors"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("input out of range, must be between 1 and 3999")
	}

	numerals := []RomanNumeral{
		{Value: 1000, Symbol: "M"},
		{Value: 900, Symbol: "CM"},
		{Value: 500, Symbol: "D"},
		{Value: 400, Symbol: "CD"},
		{Value: 100, Symbol: "C"},
		{Value: 90, Symbol: "XC"},
		{Value: 50, Symbol: "L"},
		{Value: 40, Symbol: "XL"},
		{Value: 10, Symbol: "X"},
		{Value: 9, Symbol: "IX"},
		{Value: 5, Symbol: "V"},
		{Value: 4, Symbol: "IV"},
		{Value: 1, Symbol: "I"},
	}

	result := ""

	for _, numeral := range numerals {
		for input >= numeral.Value {
			result += numeral.Symbol
			input -= numeral.Value
		}
	}

	return result, nil
}
