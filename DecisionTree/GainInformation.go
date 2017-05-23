package DecisionTree

import "fmt"

func GainInformation(features map[string][]int, targetFeature string, desorder func([]int) float64, mapGroups func(map[string][]int) map[string][]int) map[string]float64 {
	var gi = make(map[string]float64)

	for key, value := range features {

		var total = float64(len(value))
		if key != targetFeature {

			var subGroup = make(map[int][]int)
			for i, v := range value {
				subGroup[v] = append(subGroup[v], features[targetFeature][i])
			}

			fmt.Println("\nFT:", key, "")
			var sumDesorders float64
			for k, v := range subGroup {
				fmt.Println("P(", k, ") | ", v, " = ", (float64(len(v)) / total))
				fmt.Println("D(", k, ") = ", desorder(v))
				sumDesorders += (float64(len(v)) / total) * desorder(v)
			}

			fmt.Println("D(", features[targetFeature], ") = ", desorder(features[targetFeature]))
			gi[key] = desorder(features[targetFeature]) - sumDesorders

		}
	}

	return gi
}
