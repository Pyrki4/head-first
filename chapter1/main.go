package main

import "fmt"

func main() {
	var price = 100
	fmt.Println("Price is", price, "dollars.")
	var taxRate = 0.08
	var tax = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")
	var total = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")
	var availableFunds = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))
}
