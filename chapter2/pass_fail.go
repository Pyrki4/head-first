package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	grade, err := strconv.ParseInt(input, 0, 32)

	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}
	//if grade == 100 {
	//	status = "Perfect!"
	//} else if grade >= 60 {
	//	status = "You pass."
	//} else {
	//	status = "You fail!"
	//}

	fmt.Println(status)
}
