package main

import (
	"fmt"
	"golang/chapter4/average"
)

func main() {
	arr := []float64{1.1, 2.2, 3.3, 4.4}
	fmt.Println(average.AvgFloatArray(arr...))
	fmt.Println(average.AvgFloatArray(5.5, 6.6, 7.7, 8.8))
}
