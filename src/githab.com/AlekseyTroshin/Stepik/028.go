package main

import (
	"fmt"
)

var MINUT = 60
var HOUR  = MINUT * 60

func main() {
	var a int
	fmt.Scan(&a)
	hours := a / HOUR
	minuts := (a - hours * HOUR) / MINUT
	fmt.Printf("It is %d hours %d minutes.", hours, minuts)
}
