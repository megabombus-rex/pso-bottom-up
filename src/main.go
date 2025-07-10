package main

import (
	"PSO/src/experiments"
	"PSO/src/problems"
	//"PSO/src/visualization"
	//"os"
)

func main() {
	//fmt.Println("Hello world.")

	//min_max_positions := []swarm.MinMaxPair{
	//	{Min: 0.0, Max: 1.0},
	//	{Min: 0.0, Max: 1.0},
	//}

	//min_max_velocities := []swarm.MinMaxPair{
	//		{Min: -5.0, Max: 5.0},
	//		{Min: -5.0, Max: 5.0},
	//}

	//particle, err := swarm.CreateRandomParticle(1, 2, min_max_positions, min_max_velocities)

	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//fmt.Printf("%+v\n", *particle)

	//initial_swarm := swarm.CreateInitialSwarm_GBest(10, 2, min_max_positions, min_max_velocities, 0.5, 0.5, 0.5)
	//fmt.Printf("%+v\n", *initial_swarm)

	//os.MkdirAll("output", 0755)

	//particlePositions := [][]float64{
	//		{2.5, 3.1},  // Particle 1
	//		{-1.2, 0.8}, // Particle 2
	//		{0.0, 0.0},  // Global minimum
	//		{4.1, -2.3}, // Particle 4
	//		{-3.2, 1.5}, // Particle 5
	//}

	//heatmap := visualization.CreateHeatmapRastrigin(particlePositions)
	//f1, _ := os.Create("output/rastrigin_heatmap.html")
	//heatmap.Render(f1)
	//f1.Close()

	//experiments.RastriginTestPBest(20, 2, 50)
	experiments.RastriginTestPBest(30, 2, 15, problems.Sphere_fitness, -5.12, 5.12)
	experiments.RastriginTestPBest(30, 2, 15, problems.Rastrigin_fitness, -5.12, 5.12)
}
