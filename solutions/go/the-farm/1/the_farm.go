package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	numberOfCows int
	details      string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.details)
}

func DivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	totalFodder, err := fodderCalculator.FodderAmount(numberOfCows)

	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fodderCalculator.FatteningFactor()

	if err != nil {
		return 0, err
	}

	return totalFodder * fatteningFactor / float64(numberOfCows), nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows > 0 {
		return DivideFood(fodderCalculator, numberOfCows)
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numberOfCows int) error {
	switch {
	case numberOfCows < 0:
		return &InvalidCowsError{numberOfCows: numberOfCows, details: "there are no negative cows"}
	case numberOfCows == 0:
		return &InvalidCowsError{numberOfCows: numberOfCows, details: "no cows don't need food"}
	default:
		return nil
	}
}
