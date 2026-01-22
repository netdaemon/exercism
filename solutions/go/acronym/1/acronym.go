package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	r := regexp.MustCompile(`[a-zA-Z|']+`)
	words := r.FindAllString(s, -1)

	var sb strings.Builder

	for _, word := range words {
		if strings.TrimSpace(word) == "" {
			continue
		}
		sb.WriteString(strings.ToUpper(string(word[0])))
	}

	return sb.String()
}
