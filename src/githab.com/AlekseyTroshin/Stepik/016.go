package main

import (
	"fmt"
)

func main() {
	var res, a, b string
	fmt.Scan(&a)
	fmt.Scan(&b)
	count := 0

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] && count == 0 {
				res += a[i] + " "
			}
		}
	}

	fmt.Println(res)
}
