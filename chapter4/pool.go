package main

import "fmt"

func main() {
	numbers := []int{3, 16, -2, 10, 23, 12}

	for i, num := range numbers {
		if num >= 10 && num <= 20 {
			fmt.Println(i, num)
		}
	}
}
