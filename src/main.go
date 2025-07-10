package main

import (
	"PSO/src/experiments"
	"PSO/src/problems"
)

func main() {
	experiments.ExperimentPSOBest(30, 2, 15, problems.Sphere_fitness, -5.12, 5.12)
	experiments.ExperimentPSOBest(30, 2, 15, problems.Rastrigin_fitness, -5.12, 5.12)
}
