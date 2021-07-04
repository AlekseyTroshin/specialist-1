package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	fmt.Scan(&a)
    
    if len(a) < 10 || strings.Contains(a, "@") {
    	fmt.Println("Некорректный логин")
    	return
    }

	fmt.Scan(&b)
    if !strings.Contains(b, "@") || !strings.Contains(b, ".") {
    	fmt.Println("Некорректная почта")
    	return
    }

    fmt.Println("OK")
}