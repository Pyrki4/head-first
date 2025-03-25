package main

import (
	"fmt"
	"golang/chapter10/calendar"
	"log"
)

func main() {
	date := calendar.Date{}
	if err := date.SetYear(2025); err != nil {
		log.Fatal(err)
	}
	if err := date.SetMonth(12); err != nil {
		log.Fatal(err)
	}
	if err := date.SetDay(31); err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.Year(), date.Month(), date.Day())
}
