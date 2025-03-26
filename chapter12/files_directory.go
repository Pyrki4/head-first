package main

import (
	"fmt"
	"os"
)

func countTab(counter int) string {
	var s string
	for counter > 0 {
		s += " "
		counter--
	}
	return s
}

func RangeDirectory(dirName string, counter int) {
	dir, err := os.ReadDir(dirName)
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		if file.IsDir() {
			fmt.Println(countTab(counter), "Directory:", file.Name())
			counter++
				RangeDirectory(dirName + "/" + file.Name(), counter)
			counter--
		} else {
			fmt.Println(countTab(counter), "File:", file.Name())
		}
	}
}

func main() {
	var count int
	RangeDirectory("my_directory", count)
}
