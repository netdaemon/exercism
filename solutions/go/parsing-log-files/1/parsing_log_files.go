package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-=~*]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)

	count := 0

	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	newLines := []string{}

	for _, line := range lines {
		found := re.FindStringSubmatch(line)

		if found != nil {
			line = fmt.Sprintf("[USR] %s %s", found[1], line)
		}

		newLines = append(newLines, line)
	}

	return newLines
}
