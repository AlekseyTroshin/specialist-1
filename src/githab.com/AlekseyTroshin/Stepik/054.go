package main

import (
	"fmt"
	"strings"
)

func check(r rune) rune {
	fmt.Println(string(r))
	return r
}

func ExampleFirstClassFunctionArgument() {
	src := "HELLO"

	strings.Map(check, src)
}

func main() {
	ExampleFirstClassFunctionArgument()
}
