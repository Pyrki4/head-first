package main

import (
	"fmt"
	"golang/chapter4/average"
	"golang/chapter6/arguments"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	arr, err := arguments.GetSliceOfFloatFromArguments(args)
	if err != nil {
		log.Fatal(err)
	}

	avg := average.AvgFloatArray(arr...)
	fmt.Println("Average:", avg)
}
