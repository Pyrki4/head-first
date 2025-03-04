package main

import (
	"fmt"
	"log"
)

func main() {
	var amount, total float64
	amount, err := paintNeeded(5.3, 4.2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	amount, err = paintNeeded(4.2, 6.2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %.2f liters\n", total)
}

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	area := width * height
	return area / 10, nil
}
