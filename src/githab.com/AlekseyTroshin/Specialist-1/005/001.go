package main

import (
	"fmt"
)

func main() {
	var a int

	for {
		fmt.Scan(&a)

		if a == 0 {
			return
		}

		fmt.Println(a)
	}
}