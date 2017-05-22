package main

import (
	"fmt"
	"math"
)

func main() {
	var x [3][5]int = [3][5]int{{1, 2, 3, 4, 19}, {2, 3, 4, 5, 6}, {3, 4, 5, 6, 7}}

	var soma float64 = 0
	for _, value := range x {
		for _, valuej := range value {
			// fmt.Print("[", i, j, "]= ", valuej, "| ")
			soma += float64(valuej)
		}
	}

	var xw = [6]int{2, 3, 45, 100, 2, 4}

	fmt.Println(variance(xw[:]))
	fmt.Println(entropy(xw[:]))
}

func variance(x []int) float64 {
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

func entropy(x []int) float64 {
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
