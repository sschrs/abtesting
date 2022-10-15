package utils

import "sort"

// Sort function sorts a []float64 array ascending or descending and return sorted array back
func Sort(data []float64, ascending ...bool) []float64 {
	var asc bool
	if len(ascending) == 0 {
		asc = false
	} else {
		asc = ascending[0]
	}

	if asc {
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
	} else {
		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})
	}
	return data
}
