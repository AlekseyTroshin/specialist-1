package main

import (
	"fmt"
)

func main() {
	a, b := vote(1, 2, 3, 4, 5)
	fmt.Println(a, b)
}

func vote(n ...int) (int, int) {
	var sum int
	for _, i := range n {
		sum += i
	}

	return len(n), sum
}
