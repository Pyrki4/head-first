package main

import "fmt"

type student struct {
	name  string
	grade float64
}

func printInfo(s student) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: 	%0.1f\n", s.grade)
}

func main() {
	s := student{"Alonzo Cole", 93.2}

	printInfo(s)
}
