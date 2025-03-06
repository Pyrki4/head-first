package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) {
	var arr []float64

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return arr, err
		}
		arr = append(arr, num)
	}
	if err = file.Close(); err != nil {
		return arr, err

	}

	if scanner.Err() != nil {
		return arr, err
	}

	return arr, nil
}

//
//file, err := os.Open("data.txt")
//if err != nil {
//log.Fatal(err)
//}
//
//scanner := bufio.NewScanner(file)
//var arr []float64
//for i := 0; scanner.Scan(); i++ {
//arr[i], err = strconv.ParseFloat(scanner.Text(), 64)
//}
//
//err = file.Close()
//if err != nil {
//log.Fatal(err)
//}
//
//if scanner.Err() != nil {
//log.Fatal(scanner.Err())
//}
//
//fmt.Println(arr)
