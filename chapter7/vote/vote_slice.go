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
	var names []string
	var votes []int

	for _, line := range lines {
		var matched bool
		for i, name := range names {
			if line == name {
				votes[i] += 1
				matched = true
			}
		}

		if !matched {
			names = append(names, line)
			votes = append(votes, 1)
		}

	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, votes[i])
	}
}
