package main

import (
	"fmt"
	"golang/chapter3/divide"
)

func main() {
	n, err := divide.Divide(5.2, 1.2)
	fmt.Println(n, err)
}
