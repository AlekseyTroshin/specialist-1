package main

import "fmt"

func main() {
	var a, b int
    fmt.Scan(&a, &b)
    
    fmt.Println("Периметр:", a * 2 + b * 2)
    fmt.Println("Площадь:", a * b)
}