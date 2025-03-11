package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) {
	var arr []float64

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		arr = append(arr, num)
	}

	if file.Close() != nil {
		return nil, err

	}

	if scanner.Err() != nil {
		return nil, err
	}

	return arr, nil
}
