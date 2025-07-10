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

func CreateRandomParticle(id int, dimensions int, minMaxPositions []MinMaxPair, minMaxVelocities []MinMaxPair) (*Particle, error) {
	if len(minMaxPositions) != dimensions || len(minMaxVelocities) != dimensions {
		return nil, fmt.Errorf("empty data arrays")
	}

	positions := make([]float64, dimensions)
	velocities := make([]float64, dimensions)

	for i := range dimensions {
		positions[i] = random_extensions.RandRangeFloat64(minMaxPositions[i].Min, minMaxPositions[i].Max)
		velocities[i] = random_extensions.RandRangeFloat64(minMaxVelocities[i].Min, minMaxVelocities[i].Max)
	}

	particle := Particle{Id: id, Dimensions: dimensions, Positions: positions, Velocities: velocities}
	fmt.Println("Particle created: ", particle)

	return &particle, nil
}
