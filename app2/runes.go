package main

import "fmt"

func main() {
	asciiString := "ABCDE"
	utf8String := "БГДЖИ"

	asciiBytes := []rune(asciiString)
	utf8Bytes := []rune(utf8String)

	asciiBytesPartial := asciiBytes[:3]
	utf8BytesPartial := utf8Bytes[3:]

	fmt.Println(string(asciiBytesPartial))
	fmt.Println(string(utf8BytesPartial))
}
