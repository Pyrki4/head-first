package main

import "fmt"

func goroutine(c chan string) {
	c <- "a"
	c <- "b"
	c <- "c"
}

func main() {
	c := make(chan string, 3)
	go goroutine(c)
	for i:= 0; i < cap(c); i++ {
		fmt.Println(<- c)
	}
}