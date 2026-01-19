package atbash

import "strings"

func Atbash(s string) string {
	atbashed := ""

	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			atbashed += string('z' - r + 'A')
		case r >= 'a' && r <= 'z':
			atbashed += string('z' - r + 'a')
		case r >= '0' && r <= '9':
			atbashed += string(r)
		default:
			continue
		}

		if len(atbashed)%6 == 5 {
			atbashed += " "
		}
	}

	return strings.TrimSpace(atbashed)
}
