package main

import "fmt"

func snack() {
	defer fmt.Println("Closing refrigerator")
	fmt.Println("Opening refrigerator")
	panic("refrigerator is empty")
}

func main() {
 snack()
}
