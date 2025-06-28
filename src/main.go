package main

import (
	"PSO/src/swarm"
	"fmt"
)

func main() {
	fmt.Println("Hello world.")

	min_max_positions := []swarm.MinMaxPair{
		{Min: 0.0, Max: 1.0},
		{Min: 0.0, Max: 1.0},
	}

	min_max_velocities := []swarm.MinMaxPair{
		{Min: -5.0, Max: 5.0},
		{Min: -5.0, Max: 5.0},
	}

	particle, err := swarm.CreateRandomParticle(1, 2, min_max_positions, min_max_velocities)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", *particle)

	initial_swarm := swarm.CreateInitialSwarm_GBest(10, 2, min_max_positions, min_max_velocities)
	fmt.Printf("%+v\n", *initial_swarm)
}
