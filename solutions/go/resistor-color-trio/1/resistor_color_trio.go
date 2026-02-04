package resistorcolortrio

import (
	"fmt"
	"math"
)

var colorCodes = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	firstColor := colorCodes[colors[0]]
	secondColor := colorCodes[colors[1]]
	zeroCount := colorCodes[colors[2]]

	value := firstColor*10 + secondColor
	multiplier := int(math.Pow10(zeroCount))
	total := value * multiplier

	unit := ""

	switch {
	case total < 1_000:
		unit = "ohms"
	case total < 1_000_000:
		unit = "kiloohms"
		total /= 1_000
	case total < 1_000_000_000:
		unit = "megaohms"
		total /= 1_000_000
	default:
		unit = "gigaohms"
		total /= 1_000_000_000
	}

	return fmt.Sprintf("%d %s", total, unit)
}
