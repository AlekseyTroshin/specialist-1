package main

import (
	"fmt"
)

func main() {
	var a, b, res string
	fmt.Scan(&a)
	fmt.Scan(&b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[0] {
			res += string(a[i])
		}
	}

	fmt.Println(res)
}
