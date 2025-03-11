package arguments

import (
	"strconv"
)

func GetSliceOfFloatFromArguments(arguments []string) ([]float64, error) {
	var arr []float64

	for _, argument := range arguments {
		float, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			return nil, err
		}
		arr = append(arr, float)
	}
	return arr, nil
}
