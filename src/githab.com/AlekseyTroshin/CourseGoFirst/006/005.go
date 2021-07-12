package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 20, 30, 40, 50}
	arr2 := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(arr, len(arr), cap(arr))
	arr = append(arr, arr2...)

	fmt.Println(arr, len(arr), cap(arr))
}
