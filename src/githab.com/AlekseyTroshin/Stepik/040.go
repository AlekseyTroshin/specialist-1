package main

import (
	"fmt"
)

func main() {
	fmt.Println(vote(1))
	fmt.Println(vote(2))
	fmt.Println(vote(3))
	fmt.Println(vote(4))
	fmt.Println(vote(5))
}

func vote(n int) int {
	a, b := 0, 1

	for i:= 0; i < n - 1; i++ {
		a, b = b, a + b
	}

	return b
}




