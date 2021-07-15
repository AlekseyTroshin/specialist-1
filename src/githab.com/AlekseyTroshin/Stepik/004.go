package main

import "fmt"

func main() {
	var d, h, m, oneM int
	oneM = 2
	fmt.Scan(&d)
	if 0 < d && d < 360 {
		h = oneM * d / 60
		m = oneM*d - h*60
	}
	fmt.Printf("It is %d hours %d minutes.\n", h, m)
}
