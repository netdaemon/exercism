package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}

	if span < 0 {
		return 0, errors.New("span must not be negative")
	}

	numbers := make([]int64, len(digits))

	for i, r := range digits {
		num, err := strconv.Atoi(string(r))

		if err != nil {
			return 0, errors.New("digits input must only contain digits")
		}

		numbers[i] = int64(num)
	}

	var largest int64

	for i := 0; i < len(numbers)-span+1; i++ {
		product := int64(1)

		for _, d := range numbers[i : i+span] {
			product *= d
		}

		if product > largest {
			largest = product
		}
	}

	return largest, nil
}
