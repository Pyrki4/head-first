package main

import (
	"fmt"
	"golang/chapter5/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	votes := make(map[string]int)

	for _, line := range lines {
		votes[line]++
	}
	for name, count := range votes {
		fmt.Printf("%s : %d\n", name, count)
	}
}
