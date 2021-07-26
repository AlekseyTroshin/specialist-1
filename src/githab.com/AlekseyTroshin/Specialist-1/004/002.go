package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	var r rune
	r = 555
	rr := "HELLOOOO"
	fmt.Println("Rune ", r)
	fmt.Printf("Rune %c\n", r)
	fmt.Println("Rune - 1", utf8.RuneCountInString(rr))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println("aaa" < "bbb")
}