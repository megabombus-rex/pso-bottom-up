package random_extensions

import "math/rand/v2"

func RandRangeInteger(min, max int) int {
	return rand.IntN(max-min) + min
}

func RandRangeFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
