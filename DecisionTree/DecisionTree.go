package DecisionTree

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type DecisionTree struct {
	Target string
}

func (d *DecisionTree) ReadCSV(filename string) (map[string][]int, error) {

	file, e := os.Open("test.csv")
	if e != nil {
		return nil, e
	}
	lines, er := csv.NewReader(file).ReadAll()
	file.Close()
	if er != nil {
		return nil, er
	}

	var MapResult map[string][]int = make(map[string][]int)

	for i, val := range lines {
		if i > 0 {
			for j, valj := range val {
				v, erAtois := strconv.Atoi(valj)
				if erAtois != nil {
					return nil, erAtois
				}
				MapResult[lines[0][j]] = append(MapResult[lines[0][j]], v)
			}
		}

	}

	fmt.Println(MapResult)

	return MapResult, nil
}

func (d *DecisionTree) SplitDataset(DataSet map[string][]int, keyToSplit string, mapDataSet func([]int) []int) map[int]map[string][]int {
	var DataSetTansformed map[string][]int = DataSet
	var mapSplit []int = DataSet[keyToSplit]

	delete(DataSetTansformed, keyToSplit)
	fmt.Println(mapSplit)
	fmt.Println(DataSetTansformed)

	var result = make(map[int]map[string][]int, 0)

	for key, value := range DataSetTansformed {
		fmt.Println(key)
		for i, x := range value {
			var klasse int = mapSplit[i]

			if result[klasse] == nil {
				result[klasse] = make(map[string][]int)
			}

			if result[klasse][key] == nil {
				result[klasse][key] = make([]int, 0)
			}

			result[klasse][key] = append(result[klasse][key], x)
		}

	}
	return result

}
