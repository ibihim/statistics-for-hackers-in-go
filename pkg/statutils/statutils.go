package statutils

import (
	"math/rand"
)

// Shuffle shuffles given set
func Shuffle(set []float64) {
	rand.Shuffle(len(set), func(i, j int) {
		set[i], set[j] = set[j], set[i]
	})
}
