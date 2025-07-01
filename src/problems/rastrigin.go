package problems

import "math"

func Rastrigin_fitness(dimensions int, position []float64) float64 {
	value := float64(10 * dimensions)

	for i := range dimensions {
		value += (position[i]*position[i] - 10*math.Cos(2*position[i]*math.Pi))
	}

	return value
}
