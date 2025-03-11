package main

import (
	"fmt"
	"golang/chapter5/datafile"
	"log"
)

func main() {
	fileName := "data.txt"

	file, err := datafile.GetFloats(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)

}
