package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c string
	var numA, numB, numC int
	fmt.Scan(&a, &b, &c)

	numA, _ = strconv.Atoi(a)
	numB, _ = strconv.Atoi(b)
	numC, _ = strconv.Atoi(c)

	if (a == "один" || a == "раз") && b == "два" && c == "три" || 
		numA == 1 && numB == 2 && numC == 3 {
		fmt.Println("ОК")
	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}
}