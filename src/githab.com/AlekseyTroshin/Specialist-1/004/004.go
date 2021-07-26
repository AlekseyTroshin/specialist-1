package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c)
	
	d = math.Pow(b, 2)-4*a*c

	if d < 0 || a == 0 || b == 0 {
		fmt.Println("корней нет")
	} else if d == 0 {
		fmt.Println("один корень")
	} else {
		fmt.Println("два корня")
	}
}