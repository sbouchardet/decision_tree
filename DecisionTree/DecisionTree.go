package DecisionTree

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ReadCSV(filename string) (map[string][]int, error) {

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
