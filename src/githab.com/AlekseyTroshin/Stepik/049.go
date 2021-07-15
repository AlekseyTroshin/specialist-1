package main

import (
	"fmt"
	"unicode"
)

func main() {
	var ch string
	fmt.Scan(&ch)
	chRune := []rune(ch)

	for _, a := range chRune {
		if !unicode.Is(unicode.Latin, a) && !(a >= 48 && a <= 57) || len(chRune) < 5 {
			fmt.Println("Wrong password")
			return
		}
	}

	fmt.Println("Ok")
}
