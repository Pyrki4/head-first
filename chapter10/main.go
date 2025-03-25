package main

import (
	"fmt"
	"golang/chapter10/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("Birthday")
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(2004)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(8)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(20)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event)
}
