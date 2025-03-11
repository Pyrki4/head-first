package main

import "fmt"

type Part struct {
	description string
	count       int
}

func doublePart(p *Part) {
	p.count *= 2
}

func main() {
	fuses := Part{description: "Fuses", count: 5}
	doublePart(&fuses)
	fmt.Println(fuses)
}
