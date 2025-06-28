package swarm

type GBestSwarm struct {
	size          int
	global_best_p Particle
	particles     []Particle
}

func CreateInitialSwarm_GBest(size int, dimensions int, min_max_positions []MinMaxPair, min_max_velocities []MinMaxPair) *GBestSwarm {

}
