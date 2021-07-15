package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, res string
	fmt.Scan(&a)

	for _, i := range a {
		if strings.Count(a, string(i)) == 1 {
			res += string(i)
		}
	}

	fmt.Println(res)
}
