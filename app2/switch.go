package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize() {
	switch rand.Intn(3) + 1 {
	case 1:
		fmt.Println("You win a cruise")
	case 2:
		fmt.Println("You win a car")
	case 3:
		fmt.Println("You win a goat")
	default:
		panic("invalid door number")
	}
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	rand.New(src)
	awardPrize()
}
