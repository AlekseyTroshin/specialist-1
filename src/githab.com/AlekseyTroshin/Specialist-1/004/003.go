package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a float64
	var s string
	fmt.Scan(&s)
	byteStr := utf8.RuneCountInString(s)
	a = float64(byteStr) * 0.23
    
    if a > 1 {
    	fmt.Printf("%d р. %d коп.\n", int(a), int((a - float64(int(a))) * 100))
    } else {
    	fmt.Printf("%d коп.\n", int((a - float64(int(a))) * 100))
    }
}