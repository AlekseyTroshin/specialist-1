package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, xa, xb float64
	fmt.Scan(&a, &b, &c)
	
	d = math.Pow(b, 2)-4*a*c

	if d >= 0 && b != 0 {
		d = math.Sqrt(d)
		xa = (-b + d) / (2 * a)
		xb = (-b - d) / (2 * a)
		if xa == xb || d == 0 {
			fmt.Println("один корень")
		} else {
			fmt.Println("два корня")
		}
        return
	} 
    
    fmt.Println("корней нет")
}