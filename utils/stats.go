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
func Variance(data []float64, sample ...bool) float64 {
	var minus int
	if len(sample) <= 0 {
		minus = 1
	} else {
		if sample[0] == false {
			minus = 0
		} else {
			minus = 1
		}
	}
	mean := Mean(data)
	sum := 0.0
	for _, i := range data {
		sum += math.Pow(i-mean, 2)
	}
	return sum / float64(len(data)-minus)
}

// Std returns the standard deviation of a []float array
func Std(data []float64, sample ...bool) float64 {
	return math.Sqrt(Variance(data, sample...))
}

// Max returns the maximum value of an array
func Max(data []float64) float64 {
	return Sort(data, false)[0]
}

// Min returns the minimum value of an array
func Min(data []float64) float64 {
	return Sort(data, true)[0]
}
