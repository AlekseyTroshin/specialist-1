package main

import (
	"fmt"
)


func main() {
	var a, b, res int
	fmt.Scan(&a)
	fmt.Scan(&b)
	for ; a <= b; b-- {
		if a == b {
			fmt.Println("NO")
			break
		}
		if b % 7 == 0 {
			res = b
			fmt.Println(res)
			break
		}
	}	
}





