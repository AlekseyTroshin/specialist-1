package main

import (
	"fmt"
)

func main() {
	a, b, c := 0, 1, 0
	var enter int
	fmt.Scan(&enter)
	for i := 0; ; i++ {
		if i < 2 {
			continue
		}

		c = a + b
		a = b
		b = c

		if enter < c {
			fmt.Println(-1)
			break
		} else if enter == c {
			fmt.Println(i)
			break
		}
	}
}
