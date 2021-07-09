package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	arrNums := []int{}
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		fmt.Scan(&b)
		arrNums = append(arrNums, b)
	}
	fmt.Scan(&c, &d)
	fmt.Println(arrNums[c] + arrNums[d])
}
