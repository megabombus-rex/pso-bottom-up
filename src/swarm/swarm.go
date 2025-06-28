package swarm

type GBestSwarm struct {
	size          int
	global_best_p Particle
	particles     []Particle
}
