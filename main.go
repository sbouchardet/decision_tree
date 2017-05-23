package main

import (
	. "decision_tree/DecisionTree"
	"fmt"
)

func main() {
	// var x [3][5]int = [3][5]int{{1, 2, 3, 4, 19}, {2, 3, 4, 5, 6}, {3, 4, 5, 6, 7}}
	//
	// var soma float64 = 0
	// for _, value := range x {
	// 	for _, valuej := range value {
	// 		// fmt.Print("[", i, j, "]= ", valuej, "| ")
	// 		soma += float64(valuej)
	// 	}
	// }

	// var xw = [4]int{2, 2, 2, 3}

	// var features = make(map[string][]int)
	// features["ft1"] = []int{1, 1, 0, 0}
	// features["ft2"] = []int{0, 1, 1, 0}
	// features["target"] = []int{0, 0, 1, 1}
	// fmt.Println(features)
	// fmt.Println(GainInformation(features, "target", Entropy, func(x map[string][]int) map[string][]int { return x }))

	values, err := ReadCSV("teste.csv")
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(GainInformation(values, "tg", Entropy, func(x map[string][]int) map[string][]int { return x }))
}
