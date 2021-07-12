// задача H
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, sum int
	arrNums := []int{}
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		arrNums = append(arrNums, b)
	}
	fmt.Scan(&c, &d)
	c--
	d--
	for i, num := range arrNums {
		if c <= i && d >= i {
			sum += num
		}
	}
	fmt.Println(sum)
}