package chance

import "math/rand"
// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    result := rand.Intn(20)
	if result == 0 {
        return 1
    } else {
        return result
    }
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	sli := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
    rand.Shuffle(len(sli), func(i, j int) { sli[i], sli[j] = sli[j], sli[i]})
    return sli
}
