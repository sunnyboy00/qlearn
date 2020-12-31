package learn

import (
	"math/rand"
)

func generateRandomNumber(min, max int) int {
	return rand.Intn((max - min) + min)
}

func random() float64 {
	return rand.Float64()
}
