package DecisionTree

func GainInformation(features map[string][]int, targetFeature string, desorder func([]int) float64, mapGroups func(map[string][]int) map[string][]int) map[string]float64 {
	var gi = make(map[string]float64)

	for key, value := range features {

		var total = float64(len(value))
		if key != targetFeature {

			var subGroup = make(map[int][]int)
			for i, v := range value {
				subGroup[v] = append(subGroup[v], features[targetFeature][i])
			}

			var sumDesorders float64
			for _, v := range subGroup {
				sumDesorders += float64(len(v)) / total * desorder(v)
			}

			gi[key] = desorder(value) - sumDesorders

		}
	}

	return gi
}
