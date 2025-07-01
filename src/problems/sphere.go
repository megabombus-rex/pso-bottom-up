package problems

func Sphere_fitness(dimensions int, position []float64) float64 {
	value := float64(0.0)

	for i := range dimensions {
		value += (position[i] * position[i])
	}

	return value
}
