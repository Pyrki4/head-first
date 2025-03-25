package main

import (
	"fmt"
	"golang/chapter8/geo"
)

func main() {
	location := geo.Landmark{
		Name: "The Googleplex",
		Coordinates: geo.Coordinates{
			Latitude:  37.42,
			Longitude: -122.08},
	}
	fmt.Println(location)
}
