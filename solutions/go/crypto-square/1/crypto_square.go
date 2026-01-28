package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	r := regexp.MustCompile(`\W`)
	pt = strings.ToLower(r.ReplaceAllString(pt, ""))

	size := len(pt)
	columns := int(math.Ceil(math.Sqrt(float64(size))))
	rows := int(math.Ceil(float64(size) / float64(columns)))

	pt += strings.Repeat(" ", columns*rows-size)
	sb := strings.Builder{}

	for i := range columns {
		for j := range rows {
			sb.WriteByte(pt[i+j*columns])
		}

		if i < columns-1 {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}
