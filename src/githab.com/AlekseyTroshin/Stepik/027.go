package main

import (
	"fmt"
)

func main() {
	var a, res string
	fmt.Scan(&a)
	length := len(a)
	for i := length - 1; i >= 0; i-- {
		if string(a[i]) == "0" {
			break
		}
		res += string(a[i])
	}

	fmt.Println(res)
}
