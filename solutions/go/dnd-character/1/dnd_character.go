package dndcharacter

import (
	"math"
	"math/rand"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	var rolls []int = []int{
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
	}

	minRoll := slices.Min(rolls)
	minIndex := slices.Index(rolls, minRoll)

	score := 0

	for i, roll := range rolls {
		if i == minIndex {
			continue
		}

		score += roll
	}

	return score
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitutionScore := Ability()

	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitutionScore,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitutionScore),
	}
}
