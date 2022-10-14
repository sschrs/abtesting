package utils

import "math"

func Mean(data []float64) float64 {
	sum := 0.0
	for _, i := range data {
		sum += i
	}
	return sum / float64(len(data))
}

func Variance(data []float64) float64 {
	mean := Mean(data)
	sum := 0.0
	for _, i := range data {
		sum += math.Pow(i-mean, 2)
	}
	return sum / float64(len(data))
}

func Std(data []float64) float64 {
	return math.Sqrt(Variance(data))
}
