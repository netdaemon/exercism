package wordy

import (
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
    if match, _ := regexp.MatchString(`What is -?[0-9]+(?: (?:plus|minus|divided by|multiplied by) -?[0-9]+)*\?`, question); !match {
        return 0, false
    }
    
	reNumbers := regexp.MustCompile(`-?[0-9]+`)
	reOperators := regexp.MustCompile(`plus|minus|multiplied|divided`)

	numbers := reNumbers.FindAllString(question, -1)
	operators := reOperators.FindAllString(question, -1)

	if len(operators)+1 != len(numbers) {
		return 0, false
	}

	total, _ := strconv.Atoi(numbers[0])

	for i, op := range operators {
		n, _ := strconv.Atoi(numbers[i+1])
		switch op {
		case "plus":
			total += n
		case "minus":
			total -= n
		case "multiplied":
			total *= n
		case "divided":
			total /= n
		default:
			return 0, false
		}
	}

	return total, true
}
