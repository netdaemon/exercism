package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be 1 or above")
	}

	if n == 1 {
		return 2, nil
	}

	counter := 1

	for i := 3; i < math.MaxInt; i += 2 {
		if isPrime(i) {
			counter++
			if counter == n {
				return i, nil
			}
		}
	}

	return 0, nil
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
