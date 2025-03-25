package main

import "fmt"

type Gallons float64

func (g Gallons) ToLitters() Litters {
	return Litters(g * 3.785)
}

type Litters float64

func (l Litters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

type Milliliters float64

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m / 4546.1)
}

func (m Milliliters) ToLitters() Litters {
	return Litters(m / 1000.0)
}

func main() {
	bottle := Milliliters(1500)
	fmt.Println(bottle.ToLitters())
	fmt.Println(bottle.ToGallons())
}
