package main

import "fmt"

func main() {
	var a, b int
	arr := []int{}
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		arr = append(arr, b)
	}

	for i := range arr {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
}
