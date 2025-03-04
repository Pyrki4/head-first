package main

import "fmt"

func createPointer() *float64 {
	myFloat := 98.5
	return &myFloat
}

func printPoiner(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func main() {
	var myFloatPointer = createPointer()
	fmt.Println(*myFloatPointer)

	myBool := true
	printPoiner(&myBool)
}
