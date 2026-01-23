package bottlesong

import (
	"fmt"
	"strings"
)

var numberToWord = map[int]string{
	0:  "no",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

func Recite(startBottles, takeDown int) []string {
	var lines []string

	for takeDown > 0 {
		count := numberToWord[startBottles]
		remaining := strings.ToLower(numberToWord[startBottles-1])

		countBottlesStr := "bottles"
		remainingStr := "bottles"

		if startBottles == 1 {
			countBottlesStr = "bottle"
		}

		if startBottles-1 == 1 {
			remainingStr = "bottle"
		}

		lines = append(lines, fmt.Sprintf("%s green %s hanging on the wall,", count, countBottlesStr))
		lines = append(lines, fmt.Sprintf("%s green %s hanging on the wall,", count, countBottlesStr))
		lines = append(lines, "And if one green bottle should accidentally fall,")
		lines = append(lines, fmt.Sprintf("There'll be %s green %s hanging on the wall.", remaining, remainingStr))

		if takeDown > 1 {
			lines = append(lines, "")
		}

		takeDown--
		startBottles--
	}

	return lines
}
