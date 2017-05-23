package DecisionTree

import (
	"math"
)

func Variance(x []int) float64 {
	var sizeOfX float64 = float64(len(x))
	var sum float64 = 0
	for _, value := range x {
		sum += float64(value)
	}
	var average = sum / sizeOfX
	var sumDif float64 = 0
	for _, value := range x {
		sumDif += math.Pow(float64(value)-average, 2)
	}

	return sumDif / (sizeOfX - 1)
}

func Entropy(x []int) float64 {
	var ent float64 = 0
	var sliceOfDifX = make(map[int]int)
	for _, val := range x {
		sliceOfDifX[val] += 1
	}

	var sizeOfX = float64(len(x))
	var pOfKey = make(map[int]float64)

	for k, v := range sliceOfDifX {
		pOfKey[k] = float64(v) / sizeOfX
		ent += pOfKey[k] * math.Log2(pOfKey[k])
	}

	return -1 * ent

}
