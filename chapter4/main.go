package main

import (
	"fmt"
	"golang/chapter4/average"
)

func main() {
	arr := []float64{5.2, 4.6, 7.8}
	fmt.Println(arr)

	fmt.Println(average.AvgFloatArray(arr))
}
