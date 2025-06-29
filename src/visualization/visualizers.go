package visualization

import (
	rastrigin "PSO/src/problems"
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func generateRastriginData(xMin, xMax, yMin, yMax float64, resolution int) []opts.Chart3DData {
	var data []opts.Chart3DData

	stepX := (xMax - xMin) / float64(resolution-1)
	stepY := (yMax - yMin) / float64(resolution-1)

	for i := 0; i < resolution; i++ {
		for j := 0; j < resolution; j++ {
			x := xMin + float64(i)*stepX
			y := yMin + float64(j)*stepY
			z := rastrigin.Rastrigin_fitness(2, []float64{x, y})

			data = append(data, opts.Chart3DData{
				Value: []interface{}{x, y, z},
			})
		}
	}

	return data
}

func generateContourData(xMin, xMax, yMin, yMax float64, resolution int) ([][]float64, []float64, []float64) {
	var data [][]float64
	var xAxis, yAxis []float64

	stepX := (xMax - xMin) / float64(resolution-1)
	stepY := (yMax - yMin) / float64(resolution-1)

	for i := 0; i < resolution; i++ {
		xAxis = append(xAxis, xMin+float64(i)*stepX)
		yAxis = append(yAxis, yMin+float64(i)*stepY)
	}

	for i := 0; i < resolution; i++ {
		var row []float64
		y := yMin + float64(i)*stepY
		for j := 0; j < resolution; j++ {
			x := xMin + float64(j)*stepX
			z := rastrigin.Rastrigin_fitness(2, []float64{x, y})
			row = append(row, z)
		}
		data = append(data, row)
	}

	return data, xAxis, yAxis
}

func Create3DSurface() *charts.Surface3D {
	surface3d := charts.NewSurface3D()

	// Generate Rastrigin function data
	data := generateRastriginData(-5.12, 5.12, -5.12, 5.12, 50)

	surface3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "2D Rastrigin Function",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeWesteros,
		}),
	)

	var opacity float32 = 0.8
	surface3d.AddSeries("Rastrigin", data).
		SetSeriesOptions(
			charts.WithItemStyleOpts(opts.ItemStyle{
				Opacity: &opacity,
			}),
		)

	return surface3d
}

func CreateHeatmapRastrigin(particlePositions [][]float64) *charts.HeatMap {
	heatmap := charts.NewHeatMap()

	data, xAxis, yAxis := generateContourData(-5.12, 5.12, -5.12, 5.12, 50)

	var heatData []opts.HeatMapData

	for i, row := range data {
		for j, val := range row {
			heatData = append(heatData, opts.HeatMapData{
				Value: [3]interface{}{j, i, val},
			})
		}
	}

	var calculable bool = true

	var xLabels, yLabels []string
	for _, v := range xAxis {
		xLabels = append(xLabels, fmt.Sprintf("%.1f", v))
	}
	for _, v := range yAxis {
		yLabels = append(yLabels, fmt.Sprintf("%.1f", v))
	}
	heatmap.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "2D Rastrigin Function - Contour Plot",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Type: "category",
			Data: xLabels,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Type: "category",
			Data: yLabels,
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: &calculable,
			Max:        100,
			Min:        0,
		}),
	)
	heatmap.AddSeries("Rastrigin", heatData)

	var particleData []opts.HeatMapData
	for _, pos := range particlePositions {
		// Convert real coordinates to grid indices
		xIdx := int((pos[0] + 5.12) / 10.24 * 29)
		yIdx := int((pos[1] + 5.12) / 10.24 * 29)

		if xIdx < 0 {
			xIdx = 0
		}
		if xIdx > 29 {
			xIdx = 29
		}
		if yIdx < 0 {
			yIdx = 0
		}
		if yIdx > 29 {
			yIdx = 29
		}

		// Add particle as a high-value point to make it visible
		particleData = append(particleData, opts.HeatMapData{
			Value: [3]interface{}{xIdx, yIdx, 150},
		})
	}

	heatmap.AddSeries("Particles", particleData,
		charts.WithItemStyleOpts(opts.ItemStyle{
			Color: "red",
		}),
	)
	return heatmap
}
