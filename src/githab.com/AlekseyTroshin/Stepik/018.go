package main

import (
	"fmt"
)

func main() {
	// var a float64
	// fmt.Scan(&a)
	a := -000012.2123
	if a <= 0 {
		fmt.Printf("число %.2f не подходит\n", a)
	} else if a > 10000 {
		fmt.Printf("%e\n", a)
	} else {
		fmt.Printf("%.4f\n", a*a)
	}

}
