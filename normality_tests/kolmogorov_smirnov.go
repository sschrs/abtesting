package normality_tests

import (
	"github.com/sschrs/abtesting/utils"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

func KSTest(data []float64) float64 {
	data = utils.Sort(data, true)
	mean := utils.Mean(data)
	std := utils.Std(data)
	frequency := map[float64]float64{}
	for _, v := range data {
		_, ok := frequency[v]
		if ok {
			frequency[v] += 1
		} else {
			frequency[v] = 1
		}
	}

	var Sx []float64
	var Fx []float64
	cumulativeSum := 0.0
	for _, value := range data {
		cumulativeSum += frequency[value]
		Sx = append(Sx, cumulativeSum/float64(len(data)))
		cdf := distuv.Normal{
			Mu:    0,
			Sigma: 1,
		}.CDF((value - mean) / std)
		Fx = append(Fx, cdf)
	}

	var DPlusValues []float64
	var DMinusValues []float64

	for i := 0; i < len(data); i++ {
		DPlusValues = append(DPlusValues, math.Abs(Fx[i]-Sx[i]))
		if i == 0 {
			DMinusValues = append(DMinusValues, math.Abs(Fx[i]))
		} else {
			DMinusValues = append(DMinusValues, math.Abs(Fx[i]-Sx[i-1]))
		}
	}

	DPlus := utils.Max(DPlusValues)
	DMinus := utils.Max(DMinusValues)
	return utils.Max([]float64{DPlus, DMinus})
}
