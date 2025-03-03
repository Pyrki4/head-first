package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, errFile := os.Stat("my.txt")
	if errFile != nil {
		log.Fatal(errFile)
	}

	fmt.Println(fileInfo.Size())
}
