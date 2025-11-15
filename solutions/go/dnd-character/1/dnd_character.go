package dndcharacter

import (
	"math"
	"math/rand/v2"
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
	score -= 10
	return int(math.Floor(float64(score)/2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	var rolls []int
	res := 0
	for i := 0; i < 4; i++ {
		roll := rand.IntN(6)
		if roll == 0 {
			rolls = append(rolls, roll)
		} else {
			rolls = append(rolls, roll)
		}
	}
	slices.Sort(rolls)
	for v := range rolls[1:] {
		res += v
	}
	return res
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	hitpoints := 10+Modifier(constitution)
	return Character{Strength: Ability(), Dexterity: Ability(), Constitution: constitution, Intelligence: Ability(), Wisdom: Ability(), Charisma: Ability(), Hitpoints: hitpoints}
}
