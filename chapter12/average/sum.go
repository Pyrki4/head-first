package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening file")
	return os.Open(fileName)
}
func CloseFile(f *os.File) {
	fmt.Println("Closing file")
	f.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	defer CloseFile(file)
	if scanner.Err() != nil {
		return nil, err
	}
	return numbers, err
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var avg float64
	for _, number := range numbers {
		avg += number
	}

	fmt.Println(avg)
}
