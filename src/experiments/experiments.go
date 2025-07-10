package experiments

import (
	"PSO/src/problems"
	"PSO/src/swarm"
	"PSO/src/visualization"
	"fmt"
	"math/rand"
	"os"
)

// based on the paper
func RastriginTestPBest(
	particleCount int, dimensions int, iterationCount int,
	fitnessFunction problems.FitnessFunctionPositional,
	minBoundary float64, maxBoundary float64) {

	function_name := problems.GetCleanFunctionName(fitnessFunction)

	// swarm parameters
	minMaxPositions := []swarm.MinMaxPair{
		{Min: -5.12, Max: 5.12},
		{Min: -5.12, Max: 5.12},
	}

	minMaxVelocities := []swarm.MinMaxPair{
		{Min: -0.005, Max: 0.005},
		{Min: -0.005, Max: 0.005},
	}

	personalLearningRate := 0.2
	globalLearningRate := 0.01
	weightInertia := 0.3

	// 1. Initialize an array of particles with random positions and velocities on D dimensions

	initialSwarm := swarm.CreateInitialSwarm_GBest(particleCount, dimensions, minMaxPositions, minMaxVelocities, personalLearningRate, globalLearningRate, weightInertia)
	fmt.Printf("%+v\n", *initialSwarm)

	previousFitnesses := make([]float64, initialSwarm.Size)
	newFitnesses := make([]float64, initialSwarm.Size)

	// 2. Evaluate the desired minimization function in D variables
	// initializing fitness
	for i, particle := range initialSwarm.Particles {
		previousFitnesses[i] = fitnessFunction(dimensions, particle.Particle.Positions)
	}

	for iteration := range iterationCount {
		fmt.Println("Iteration ", iteration, ".")
		r1Coef, r2Coef := rand.Float64(), rand.Float64()

		newVelocities := swarm.CalculateNextVelocities(initialSwarm, r1Coef, r2Coef)
		newPositions := swarm.CalculateNextPositions(initialSwarm, r1Coef, r2Coef, newVelocities)

		for i := range newPositions {
			problems.ClampPositionsBounceVelocities(newPositions[i], newVelocities[i], minBoundary, maxBoundary)
		}

		swarm.UpdateSwarmData(initialSwarm, newPositions, newVelocities)

		// new fitness
		for i, particle := range initialSwarm.Particles {
			newFitnesses[i] = fitnessFunction(dimensions, particle.Particle.Positions)
			//fmt.Println("Fitness:", new_fitnesses[i], " for particle", particle.Particle.Id, " in position: ", particle.Particle.Positions)
		}

		// 3.  Compare evaluation with particle’s previous best value (PBEST[]):
		// MINIMIZING - If current value < PBEST[] then PBEST[] = current value and PBESTx[][d] = current position in D- dimensional hyperspace
		// MAXIMIZING - If current value > PBEST[] then PBEST[] = current value and PBESTx[][d] = current position in D- dimensional hyperspace
		// here we maximize
		for i, particle := range initialSwarm.Particles {
			if newFitnesses[i] > previousFitnesses[i] {
				particle.Best_value = newFitnesses[i]
				fmt.Println("New fitness value for particle: ", particle.Particle.Id, " is ", particle.Best_value)
			}
			if newFitnesses[i] > initialSwarm.Global_best_p.Best_value {
				new_particle := swarm.Particle{
					Id:         initialSwarm.Particles[i].Particle.Id,
					Dimensions: initialSwarm.Particles[i].Particle.Dimensions,
					Positions:  initialSwarm.Particles[i].Particle.Positions,
					Velocities: initialSwarm.Particles[i].Particle.Velocities}

				initialSwarm.Global_best_p = swarm.GBestSwarmParticle{
					Particle:   new_particle,
					Best_value: newFitnesses[i]}

				fmt.Println("Found new best particle in a swarm: ", initialSwarm.Global_best_p.Particle)
			}
		}

		// replace old fitnesses with new fitnesses
		for i := range len(initialSwarm.Particles) {
			previousFitnesses[i] = newFitnesses[i]
		}

		var currentPositions [][]float64
		for _, particle := range initialSwarm.Particles {
			currentPositions = append(currentPositions, particle.Particle.Positions)
		}

		chart := visualization.CreateHeatmapFunction2D(currentPositions, fitnessFunction)

		f, _ := os.Create(fmt.Sprintf("output/pso_problem_%s_%03d.html", function_name, iteration))
		chart.Render(f)
		f.Close()

		// 4. Compare evaluation with group’s previous best (PBEST[GBEST]): If current value < PBESTCGBEST] then GBEST=particle’s array index

		// 5.Change velocity by the following formula:
		// W[dI = W[dI + ACC-CONST*rand()*(PBESTx[] [d] - PresentX[] [d]) +
		// ACC-CONST*rand()*(PBESTx[GBEST] [d] - PresentX[l[d])

		// 6. Move to PresentX[][d] + v[][d]: Loop to step 2 and repeat until a criterion is met
	}

	fmt.Println("After ", iterationCount, " iterations, best value in swarm: ", initialSwarm.Global_best_p.Best_value, " in position: ", initialSwarm.Global_best_p.Particle.Positions)
}
