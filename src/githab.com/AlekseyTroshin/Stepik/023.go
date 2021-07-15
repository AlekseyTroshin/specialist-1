package main

import "fmt"

func main() {
	array := [5]int{}
	var a, b int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	for _, i := range array {
		if i > b {
			b = i
		}
	}
	fmt.Println(b)
}
