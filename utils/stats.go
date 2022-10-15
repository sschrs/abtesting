// Package utils/stats includes some basic statistical calculations.
package utils

import (
	"math"
)

// Mean returns the average of a []float array
func Mean(data []float64) float64 {
	sum := 0.0
	for _, i := range data {
		sum += i
	}
	return sum / float64(len(data))
}

// Variance returns the variance of a []float array
func Variance(data []float64) float64 {
	mean := Mean(data)
	sum := 0.0
	for _, i := range data {
		sum += math.Pow(i-mean, 2)
	}
	return sum / float64(len(data))
}

// Std returns the standard deviation of a []float array
func Std(data []float64) float64 {
	return math.Sqrt(Variance(data))
}
