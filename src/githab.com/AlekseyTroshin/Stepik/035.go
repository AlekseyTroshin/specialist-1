package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a)
	b = 2
	for i := 1.0; ;i++{
		c = math.Pow(b, i - 1)
		if c > a {
			break
		}
		if i <= 2 {
			fmt.Print(i, " ")
			continue
		}
		fmt.Print(c, " ")
	}
}
