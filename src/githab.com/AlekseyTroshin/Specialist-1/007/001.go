package main

import (
	"fmt"
)

func main() {
	name := "Hello GO"
	fmt.Println(name)

	word := "Sample word"
	fmt.Printf("String %s\n", word)
	fmt.Printf("Bytes %s\n", word)
	rune := []rune(word)
	fmt.Println(rune)
	for _, i := range rune {
		fmt.Printf("%c ", i)
	}
	fmt.Println()
	fmt.Println(rune[3])
	fmt.Printf("%c\n", rune[3])
}
