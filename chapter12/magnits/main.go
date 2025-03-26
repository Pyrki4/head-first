package main

import (
	"fmt"
	"log"
)

func find(item string, slice []string) bool {
	for _, v := range slice {
		if item == v {
			return true
		}
	}
	return false
}
type Refrigerator []string
 func (r Refrigerator) Open() {
	 fmt.Println("Opening refrigerator")
 }
 func (r Refrigerator) Close() {
	 fmt.Println("Closing refrigerator")
 }
 func (r Refrigerator) FindFood(food string) error {
	r.Open()
	if find(food, r) {
		fmt.Println("Found", food)
	} else {
		fmt.Println("Not found", food)
	 }
	defer r.Close()
	return nil
 }

func main() {
	fridge := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk","Bananas"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
