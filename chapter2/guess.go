package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	randNumber := rand.Intn(100) + 1
	fmt.Println("Guess number at range 0 to 100.")

	var i int

	for i = 1; i <= 10; i++ {
		fmt.Print("Guess number: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		number, err := strconv.ParseUint(input, 0, 64)

		if err != nil {
			log.Fatal(err)
		}

		if uint64(randNumber) == number {
			fmt.Println("You guess the number is", randNumber)
			break
		} else if uint64(randNumber) <= number {
			fmt.Println("Oops. Your guess was HIGH.Try again")
		} else if uint64(randNumber) >= number {
			fmt.Println("Oops. Your guess was LOW.Try again")
		}
	}
	if i > 10 {
		fmt.Println("You didn’t guess my number. It was:", randNumber)
	}
}
