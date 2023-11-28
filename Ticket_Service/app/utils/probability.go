package utils

import (
	"fmt"
	"math/rand"
)

func SimulateProbability(failedPercentage int) bool {
	probability := rand.Intn(100)
	fmt.Println(probability)
	return probability >= failedPercentage
}