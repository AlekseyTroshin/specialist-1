package main

import "fmt"

func main() {
	var a, b float64
    fmt.Scan(&a, &b)  
    if int(a + b) % 2 == 0 {
    	fmt.Println("ЧЁТНОЕ")
    } else {
    	fmt.Println("НЕЧЁТНОЕ")
    }
}