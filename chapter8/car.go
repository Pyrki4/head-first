package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func main() {
	mustang := car{"Mustang Cobra", 225}
	nitroBoost(&mustang)
	fmt.Println(mustang)
}
