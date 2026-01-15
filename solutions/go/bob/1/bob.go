// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isYelling(remark string) bool {
	hasLetters := strings.IndexFunc(remark, unicode.IsLetter) >= 0

	return hasLetters && strings.ToUpper(remark) == remark
}

func isSilence(remark string) bool {
	return remark == ""
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case isYelling(remark) && !isQuestion(remark):
		return "Whoa, chill out!"
	case isYelling(remark) && isQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isSilence(remark):
		return "Fine. Be that way!"
	case isQuestion(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}
