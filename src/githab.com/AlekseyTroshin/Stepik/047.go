package main

import (
	"fmt"
)

func main() {
	var a, result string
	fmt.Scan(&a)
	for i, letter := range a {
		if i%2 != 0 {
			result += string(letter)
		}
	}

	fmt.Println(result)
}
