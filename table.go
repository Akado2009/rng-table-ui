package main

import (
	"math"
	"math/rand"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}

func randomFloat(min, max float64) float64 {
	return toFixed(min + rand.Float64() * (max - min), 1)
}

func generateRandomTable(height, width int, min, max float64) [][]float64 {
	result := make([][]float64, height, height)
	for i := 0; i < height; i++ {
		row := make([]float64, width, width)
		for j := 0; j < width; j++ {
			row[j] = randomFloat(min, max)
		}
		result[i] = row
	}
	return result
}