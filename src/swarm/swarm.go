package swarm

import "fmt"

type GBestSwarmParticle struct {
	Particle   Particle
	Best_value float64
}

type GBestSwarm struct {
	Size          int
	Global_best_p GBestSwarmParticle
	Particles     []GBestSwarmParticle
	c1_coef       float64 // personal learning coefficient
	c2_coef       float64 // global learning coefficient
	w1_coef       float64 // weight inertia coefficient
}

func CreateInitialSwarm_GBest(size int, dimensions int,
	min_max_positions []MinMaxPair,
	min_max_velocities []MinMaxPair,
	c1 float64, c2 float64, w1 float64) *GBestSwarm {
	particles := make([]GBestSwarmParticle, size)
	bad_particles_count := 0
	for i := range size {
		particle, err := CreateRandomParticle(i, dimensions, min_max_positions, min_max_velocities)
		if err != nil {
			fmt.Println(err.Error())
			bad_particles_count++
		}

		particles[i] = GBestSwarmParticle{*particle, 0.0}
	}

	return &GBestSwarm{size - bad_particles_count, particles[0], particles, c1, c2, w1}
}

func CalculateNextVelocities(swarm *GBestSwarm, r1 float64, r2 float64) [][]float64 {
	if swarm.Size < 1 {
		return nil
	}
	swarm_next_velocities := make([][]float64, len(swarm.Particles))

	for i := range swarm.Size {
		swarm_next_velocities[i] = make([]float64, swarm.Particles[i].Particle.Dimensions) // should be equal in each particle
	}

	for i, particle := range swarm.Particles {
		for j := range len(particle.Particle.Velocities) {
			weighted_v1 := swarm.w1_coef * particle.Particle.Velocities[j]
			personal_l := swarm.c1_coef * r1 * (particle.Best_value - particle.Particle.Positions[j])
			global_l := swarm.c2_coef * r2 * (swarm.Global_best_p.Best_value - particle.Particle.Positions[j])
			val := weighted_v1 + personal_l + global_l
			swarm_next_velocities[i][j] = val
		}
	}

	return swarm_next_velocities
}

func CalculateNextPositions(swarm *GBestSwarm, r1 float64, r2 float64) [][]float64 {
	swarm_next_velocities := CalculateNextVelocities(swarm, r1, r2)
	if swarm_next_velocities == nil {
		return nil
	}

	swarm_next_positions := make([][]float64, len(swarm.Particles))

	for i := range swarm.Size {
		swarm_next_positions[i] = make([]float64, swarm.Particles[i].Particle.Dimensions) // should be equal in each particle
	}

	for i, particle_velocities := range swarm_next_velocities {
		// the len() can be done higher, if the homogenity is enforced
		for j := range len(particle_velocities) {
			swarm_next_positions[i][j] = swarm.Particles[i].Particle.Positions[j] + particle_velocities[j]
		}
	}

	return swarm_next_positions
}
