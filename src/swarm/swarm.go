package swarm

import "fmt"

type GBestSwarmParticle struct {
	Particle  Particle
	BestValue float64
}

type GBestSwarm struct {
	Size        int
	GlobalBestP GBestSwarmParticle
	Particles   []GBestSwarmParticle
	c1Coef      float64 // personal learning coefficient
	c2Coef      float64 // global learning coefficient
	w1Coef      float64 // weight inertia coefficient
}

func CreateInitialSwarm_GBest(size int, dimensions int,
	minMaxPositions []MinMaxPair,
	minMaxVelocities []MinMaxPair,
	c1 float64, c2 float64, w1 float64) *GBestSwarm {
	particles := make([]GBestSwarmParticle, size)
	badParticlesCount := 0
	for i := range size {
		particle, err := CreateRandomParticle(i, dimensions, minMaxPositions, minMaxVelocities)
		if err != nil {
			fmt.Println(err.Error())
			badParticlesCount++
		}

		particles[i] = GBestSwarmParticle{*particle, 0.0}
	}

	return &GBestSwarm{size - badParticlesCount, particles[0], particles, c1, c2, w1}
}

func CalculateNextVelocities(swarm *GBestSwarm, r1 float64, r2 float64) [][]float64 {
	if swarm.Size < 1 {
		return nil
	}

	swarmNextVelocities := make([][]float64, len(swarm.Particles))

	for i := range swarm.Size {
		swarmNextVelocities[i] = make([]float64, swarm.Particles[i].Particle.Dimensions) // should be equal in each particle
	}

	for i, particle := range swarm.Particles {
		for j := range len(particle.Particle.Velocities) {
			weightedVelocity := swarm.w1Coef * particle.Particle.Velocities[j]
			personalLearning := swarm.c1Coef * r1 * (particle.BestValue - particle.Particle.Positions[j])
			globalLearning := swarm.c2Coef * r2 * (swarm.GlobalBestP.BestValue - particle.Particle.Positions[j])
			val := weightedVelocity + personalLearning + globalLearning
			fmt.Println("Velocity calculated for particle ", particle.Particle.Id, " in dimension ", j+1, " is ", val)
			swarmNextVelocities[i][j] = val
		}
	}

	return swarmNextVelocities
}

func CalculateNextPositions(swarm *GBestSwarm, r1 float64, r2 float64, swarmNextVelocities [][]float64) [][]float64 {
	swarmNextPositions := make([][]float64, len(swarm.Particles))

	for i := range swarm.Size {
		swarmNextPositions[i] = make([]float64, swarm.Particles[i].Particle.Dimensions) // should be equal in each particle
	}

	for i, particleVelocities := range swarmNextVelocities {
		// the len() can be done higher, if the homogenity is enforced
		for j := range len(particleVelocities) {
			swarmNextPositions[i][j] = swarm.Particles[i].Particle.Positions[j] + particleVelocities[j]
		}
		fmt.Println("Particle ", swarm.Particles[i].Particle.Id, " set in position: ", swarmNextPositions[i])
	}

	return swarmNextPositions
}

func UpdateSwarmData(swarm *GBestSwarm, newPositions [][]float64, newVelocities [][]float64) {
	for i, positions := range newPositions {
		fmt.Println("For particle: ", swarm.Particles[i].Particle.Id, " setting positions: ", positions)
		for j := range len(positions) {
			swarm.Particles[i].Particle.Positions[j] = newPositions[i][j]
			swarm.Particles[i].Particle.Velocities[j] = newVelocities[i][j]
		}
	}
}
