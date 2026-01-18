package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}

	subject = strings.ToLower(subject)

	sortedSubject := sortString(subject)

	for _, candidate := range candidates {
		if len(subject) != len(candidate) {
			continue
		}
		c := strings.ToLower(candidate)

		if subject == c {
			continue
		}

		sortedCandidate := sortString(c)

		if sortedSubject == sortedCandidate {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

func sortString(s string) string {
	split := strings.Split(s, "")
	slices.Sort(split)

	return strings.Join(split, "")
}
