// Package normality_tests
// Normality tests are tests used to decide whether a distribution is normally distributed.
// Some methods used for the AB Testing are based on the assumption that the distribution is normal.
// Therefore, it is important to test the normality of the distribution.
package normality_tests

import (
	"github.com/sschrs/abtesting/utils"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

/*
JarqueBera is one of the tests used to examine the normality of a distribution.
For the Normal Distribution, the measure of curvature calculated with the moments is 0 and the measure of kurtosis is 3.
The Jarque Bera test was developed based on these measures.
For more information https://en.wikipedia.org/wiki/Jarque–Bera_test
*/
func JarqueBera(data []float64) (float64, float64) {
	mean := utils.Mean(data)
	std := utils.Std(data, false)
	moment_3 := 0.0
	moment_4 := 0.0
	for _, i := range data {
		moment_3 += math.Pow((i - mean), 3)
		moment_4 += math.Pow((i - mean), 4)
	}
	moment_3 = moment_3 / float64(len(data))
	moment_4 = moment_4 / float64(len(data))

	S := moment_3 / math.Pow(std, 3)
	K := moment_4 / math.Pow(std, 4)
	testStat := float64(len(data)) * ((math.Pow(S, 2) / 6) + (math.Pow(K-3, 2))/24)
	pValue := 1 - distuv.ChiSquared{K: 2}.CDF(testStat)
	return testStat, pValue
}
