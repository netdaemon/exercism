package cipher

import (
	"strings"
	"unicode"
)

type shift struct {
	distance int
}

type vigenere struct {
	distances []int
}

func NewCaesar() Cipher {
	return shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}

	return shift{distance: distance}
}

func (c shift) Encode(input string) string {
	var sb strings.Builder

	input = strings.ToLower(input)
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		sb.WriteRune(shiftRune(r, c.distance))
	}

	return sb.String()
}

func (c shift) Decode(input string) string {
	var sb strings.Builder

	input = strings.ToLower(input)
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		sb.WriteRune(shiftRune(r, 26-c.distance))
	}

	return sb.String()
}

func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}

	distances := []int{}

	isAllA := true

	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}

		if isAllA && r != 'a' {
			isAllA = false
		}

		distances = append(distances, int(r-'a'))
	}

	if isAllA {
		return nil
	}

	return vigenere{distances: distances}
}

func (v vigenere) Encode(input string) string {
	var sb strings.Builder

	input = strings.ToLower(input)
	distancesLength := len(v.distances)
	i := 0

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		sb.WriteRune(shiftRune(r, v.distances[i%distancesLength]))
		i++
	}

	return sb.String()
}

func (v vigenere) Decode(input string) string {
	var sb strings.Builder

	input = strings.ToLower(input)
	distancesLength := len(v.distances)
	i := 0

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		sb.WriteRune(shiftRune(r, -v.distances[i%distancesLength]))
		i++
	}

	return sb.String()
}

func shiftRune(r rune, shift int) rune {
	if shift < 0 {
		shift = 26 + shift
	}

	i := (int(r-'a') + shift) % 26

	return rune('a' + i)
}
