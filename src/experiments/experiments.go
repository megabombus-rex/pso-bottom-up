package experiments

import (
	"PSO/src/problems"
	"PSO/src/swarm"
	"fmt"
	"math/rand"
)

// based on the paper
func RastriginTestPBest(particleCount int, dimensions int, iteration_count int, fitness_function problems.FitnessFunctionPositional) {
	// swarm parameters
	min_max_positions := []swarm.MinMaxPair{
		{Min: -1.0, Max: 1.0},
		{Min: -1.0, Max: 1.0},
	}

	min_max_velocities := []swarm.MinMaxPair{
		{Min: 0.1, Max: 0.25},
		{Min: 0.1, Max: 0.25},
	}

	personal_learning_rate := 0.4
	global_learning_rate := 0.8
	weight_inertia := 0.3

	// 1. Initialize an array of particles with random positions and velocities on D dimensions

	initial_swarm := swarm.CreateInitialSwarm_GBest(particleCount, dimensions, min_max_positions, min_max_velocities, personal_learning_rate, global_learning_rate, weight_inertia)
	fmt.Printf("%+v\n", *initial_swarm)

	previous_fitnesses := make([]float64, initial_swarm.Size)
	new_fitnesses := make([]float64, initial_swarm.Size)

	// 2. Evaluate the desired minimization function in D variables
	// initializing fitness
	for i, particle := range initial_swarm.Particles {
		previous_fitnesses[i] = fitness_function(dimensions, particle.Particle.Positions)
	}

	for i := range iteration_count {
		fmt.Println("Iteration ", i, ".")
		r1_coef, r2_coef := rand.Float64(), rand.Float64()

		new_velocities := swarm.CalculateNextVelocities(initial_swarm, r1_coef, r2_coef)
		new_positions := swarm.CalculateNextPositions(initial_swarm, r1_coef, r2_coef, new_velocities)

		swarm.UpdateSwarmData(initial_swarm, new_positions, new_velocities)

		// new fitness
		for i, particle := range initial_swarm.Particles {
			new_fitnesses[i] = fitness_function(dimensions, particle.Particle.Positions)
			fmt.Println("Fitness:", new_fitnesses[i], " for particle in position: ", particle.Particle.Positions)
		}

		// 3.  Compare evaluation with particle’s previous best value (PBEST[]):
		// MINIMIZING - If current value < PBEST[] then PBEST[] = current value and PBESTx[][d] = current position in D- dimensional hyperspace
		// MAXIMIZING - If current value > PBEST[] then PBEST[] = current value and PBESTx[][d] = current position in D- dimensional hyperspace
		// here we maximize
		for i, particle := range initial_swarm.Particles {
			if new_fitnesses[i] > previous_fitnesses[i] {
				particle.Best_value = new_fitnesses[i]
				fmt.Println("Particle id: ", particle.Particle.Id, " and i: ", i)
			}
			if new_fitnesses[i] > initial_swarm.Global_best_p.Best_value {
				new_particle := swarm.Particle{
					Id:         initial_swarm.Particles[i].Particle.Id,
					Dimensions: initial_swarm.Particles[i].Particle.Dimensions,
					Positions:  initial_swarm.Particles[i].Particle.Positions,
					Velocities: initial_swarm.Particles[i].Particle.Velocities}

				initial_swarm.Global_best_p = swarm.GBestSwarmParticle{
					Particle:   new_particle,
					Best_value: new_fitnesses[i]}

				fmt.Println("Found new best particle in a swarm: ", initial_swarm.Global_best_p.Particle)
			}
		}

		// replace old fitnesses with new fitnesses
		for i := range len(initial_swarm.Particles) {
			previous_fitnesses[i] = new_fitnesses[i]
		}

		// 4. Compare evaluation with group’s previous best (PBEST[GBEST]): If current value < PBESTCGBEST] then GBEST=particle’s array index

		// 5.Change velocity by the following formula:
		// W[dI = W[dI + ACC-CONST*rand()*(PBESTx[] [d] - PresentX[] [d]) +
		// ACC-CONST*rand()*(PBESTx[GBEST] [d] - PresentX[l[d])

		// 6. Move to PresentX[][d] + v[][d]: Loop to step 2 and repeat until a criterion is met
	}

	fmt.Println("After ", iteration_count, " iterations, best value in swarm: ", initial_swarm.Global_best_p.Best_value, " in position: ", initial_swarm.Global_best_p.Particle.Positions)
}
