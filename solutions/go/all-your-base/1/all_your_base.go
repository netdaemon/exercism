package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	decimalValue, err := toDecimal(inputBase, inputDigits)

	if err != nil {
		return nil, err
	}

	outputDigits := []int{}

	if decimalValue == 0 {
		outputDigits = append(outputDigits, 0)
		return outputDigits, nil
	}

	for decimalValue > 0 {
		outputDigits = append([]int{decimalValue % outputBase}, outputDigits...)
		decimalValue /= outputBase
	}

	return outputDigits, nil
}

func toDecimal(inputBase int, inputDigits []int) (int, error) {
	decimalValue := 0

	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}

		decimalValue = decimalValue*inputBase + d
	}

	return decimalValue, nil
}
