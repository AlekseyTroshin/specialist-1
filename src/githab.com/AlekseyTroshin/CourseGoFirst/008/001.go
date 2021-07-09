package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr[2] = 222
	arr[4] = 981
	fmt.Println(arr)

	arr = [5]int{
		22,
		33,
		44,
		55,
		66,
	}
	fmt.Println(arr)

	var arr2 = [...]int{}
	fmt.Println(arr2)

	words := [2][2]string{
		{"one", "two"},
		{"three", "four"},
	}

	for _, val1 := range words {
		for i, val2 := range val1 {
			fmt.Println(i+1, val2)
		}
		fmt.Println("======")
	}
}
