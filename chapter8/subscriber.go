package main

import (
	"fmt"
	. "golang/chapter8/magazine"
)

func applyDiscounter(s *Subscriber) {
	s.Rate = 4.99
}

func main() {
	s1 := Subscriber{Name: "Vlad", Rate: 2.3, Active: true}
	applyDiscounter(&s1)
	fmt.Println("Name:", s1.Name)
	fmt.Println("Rate:", s1.Rate)

	var subscriber2 Subscriber
	subscriber2.Name = "Beth Ryan"
	fmt.Println("Name:", subscriber2.Name)
}
