package main

import (
	"log"
	"os"
)

func saveString(fileName string, str string) error {
	err := os.WriteFile(fileName, []byte(str), 0600)
	return err
}

func main() {
	if err := saveString("hindi.txt", "Namaste"); err != nil {
		log.Fatal(err)
	}
}
