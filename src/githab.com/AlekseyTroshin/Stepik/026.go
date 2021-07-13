package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	sum := 0
	fmt.Scan(&a)

	for i := 0; i < len(a); i++ {
		b, _ := strconv.Atoi(string(a[i]))
		sum += b
	}

	fmt.Println(sum)
}
