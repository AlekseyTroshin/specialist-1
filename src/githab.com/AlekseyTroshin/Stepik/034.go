package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	switch {
	case a == 1 || a%10 == 1 && a%11 != 0:
		fmt.Println(a, "korova")
	case a >= 2 && a <= 4 || a%10 >= 2 && a%10 <= 4 && a > 20:
		fmt.Println(a, "korovy")
	default:
		fmt.Println(a, "korov")
	}
}
