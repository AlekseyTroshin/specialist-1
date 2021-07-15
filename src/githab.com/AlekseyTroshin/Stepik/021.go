package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)
}
