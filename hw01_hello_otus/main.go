package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	inputString := "Hello, OTUS!"

	fmt.Println(stringutil.Reverse(inputString))
}
