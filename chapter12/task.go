package main

import (
	"errors"
	"fmt"
)

func reportPanic(){
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	}
}

func main() {
	defer reportPanic()
	panic(errors.New("error"))
}
