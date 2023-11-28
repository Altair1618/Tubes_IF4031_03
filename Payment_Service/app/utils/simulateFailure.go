package utils

import "math/rand"

func SimulateFailure(percentage int) bool {
	randomNumber := rand.Intn(101)
	return randomNumber <= percentage
}
