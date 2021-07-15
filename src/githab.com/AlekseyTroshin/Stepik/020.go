package main

import (
	"fmt"
)

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[3:6]
	fmt.Println(a, b)
	b = append(b, 9, 9, 9, 10, 10, 10)
	fmt.Println(a, b)
}
