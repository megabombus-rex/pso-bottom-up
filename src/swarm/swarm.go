package swarm

import "fmt"

type GBestSwarm struct {
	Size          int
	Global_best_p Particle
	Particles     []Particle
}

func CreateInitialSwarm_GBest(size int, dimensions int, min_max_positions []MinMaxPair, min_max_velocities []MinMaxPair) *GBestSwarm {
	particles := make([]Particle, size)
	bad_particles_count := 0
	for i := range size {
		particle, err := CreateRandomParticle(i, dimensions, min_max_positions, min_max_velocities)
		if err != nil {
			fmt.Println(err.Error())
			bad_particles_count++
		}

		particles[i] = *particle
	}

	return &GBestSwarm{size - bad_particles_count, particles[0], particles}
}
