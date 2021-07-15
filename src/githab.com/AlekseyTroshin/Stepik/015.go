package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	count := 1
	x *= 100
	y *= 100

	for {
		x += (x / 100) * p
		if x > y {
			break
		}

		count++
	}
	fmt.Println(count)
}
