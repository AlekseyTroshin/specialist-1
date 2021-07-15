package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	var z string
	fmt.Scan(&z)
	a = len(z)
	for i := 0; i < a; i++ {
		j, _ := strconv.Atoi(string(z[i]))
		b += j
		if a == i+1 {
			z += strconv.Itoa(b)
			if g := strconv.Itoa(b); len(g) <= 1 {
				break
			}
			b = 0
			a = len(z)
		}
	}

	fmt.Println(b)
}
