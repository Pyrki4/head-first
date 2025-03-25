package main

import (
	"fmt"
	"golang/chapter8/geo"
	"log"
)

func main() {
	location := geo.Landmark{}
	err := location.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location)
}
