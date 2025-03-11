package datafile

import (
	"bufio"
	"log"
	"os"
)

func GetStrings(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, err
}
