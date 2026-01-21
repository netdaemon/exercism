package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	nStr := strconv.Itoa(n)

	total := 0

	for _, str := range nStr {
		num, _ := strconv.Atoi(string(str))

		total += int(math.Pow(float64(num), float64(len(nStr))))
	}

	return n == total
}
