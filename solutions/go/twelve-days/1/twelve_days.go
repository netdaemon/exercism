package twelve

import (
	"fmt"
	"strings"
)

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string{
	"twelve Drummers Drumming,",
	"eleven Pipers Piping,",
	"ten Lords-a-Leaping,",
	"nine Ladies Dancing,",
	"eight Maids-a-Milking,",
	"seven Swans-a-Swimming,",
	"six Geese-a-Laying,",
	"five Gold Rings,",
	"four Calling Birds,",
	"three French Hens,",
	"two Turtle Doves, and",
	"a Partridge in a Pear Tree.",
}

func Verse(i int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", days[i-1], strings.Join(gifts[12-i:], " "))
}

func Song() string {
	sb := strings.Builder{}

	for verse := 1; verse <= 12; verse++ {
		sb.WriteString(Verse(verse))

		if verse < 12 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
