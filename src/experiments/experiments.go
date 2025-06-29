package experiments

import (
	rastrigin "PSO/src/problems"
	"PSO/src/swarm"
	"fmt"
)

// based on the paper
func RastriginTestPBest(particleCount int, dimensions int) {
	min_max_positions := []swarm.MinMaxPair{
		{Min: -5.0, Max: 5.0},
		{Min: -5.0, Max: 5.0},
	}

	min_max_velocities := []swarm.MinMaxPair{
		{Min: 0.0, Max: 1.0},
		{Min: 0.0, Max: 1.0},
	}

	// 1. Initialize an array of particles with random positions and velocities on D dimensions,
	initial_swarm := swarm.CreateInitialSwarm_GBest(particleCount, dimensions, min_max_positions, min_max_velocities)
	fmt.Printf("%+v\n", *initial_swarm)

	fitnesses := make([]float64, initial_swarm.Size)

	// 2. Evaluate the desired minimization function in D variables
	for i, particle := range initial_swarm.Particles {
		fitnesses[i] = rastrigin.Rastrigin_fitness(dimensions, particle.Positions)
	}

	// 3.  Compare evaluation with particle’s previous best value (PBEST[]):
	// If current value < PBEST[] then PBEST[] = current value and PBESTx[][d] = current position in D- dimensional hyperspace

	// 4. Compare evaluation with group’s previous best (PBEST[GBEST]): If current value < PBESTCGBEST] then GBEST=particle’s array index

	// 5.Change velocity by the following formula:
	// W[dI = W[dI + ACC-CONST*rand()*(PBESTx[] [d] - PresentX[] [d]) +
	// ACC-CONST*rand()*(PBESTx[GBEST] [d] - PresentX[l[d])

	// 6. Move to PresentX[][d] + v[][d]: Loop to step 2 and repeat until a criterion is met
}
