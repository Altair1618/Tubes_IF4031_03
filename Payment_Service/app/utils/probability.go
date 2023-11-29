package utils

import (
	"math/rand"
)

func SimulateProbability(failedPercentage int) bool {
	probability := rand.Intn(100)
	return probability >= failedPercentage
}