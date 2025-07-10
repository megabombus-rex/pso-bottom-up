package swarm

import (
	"PSO/src/random_extensions"
	"fmt"
)

type Particle struct {
	Id         int
	Dimensions int
	Positions  []float64
	Velocities []float64
}

type MinMaxPair struct {
	Min float64
	Max float64
}

func CreateRandomParticle(id int, dimensions int, min_max_positions []MinMaxPair, min_max_velocities []MinMaxPair) (*Particle, error) {
	if len(min_max_positions) != dimensions || len(min_max_velocities) != dimensions {
		return nil, fmt.Errorf("empty data arrays")
	}

	positions := make([]float64, dimensions)
	velocities := make([]float64, dimensions)

	for i := range dimensions {
		positions[i] = random_extensions.RandRangeFloat64(min_max_positions[i].Min, min_max_positions[i].Max)
		velocities[i] = random_extensions.RandRangeFloat64(min_max_velocities[i].Min, min_max_velocities[i].Max)
	}

	particle := Particle{Id: id, Dimensions: dimensions, Positions: positions, Velocities: velocities}
	fmt.Println("Particle created: ", particle)

	return &particle, nil
}
