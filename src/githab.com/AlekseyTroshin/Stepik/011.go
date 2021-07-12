package main

import "fmt"

func main() {
    var a, b, sum int
    fmt.Scan(&a)
    
    
    for ; 0 < a; a-- {
    	fmt.Scan(&b)
        if b % 8 == 0 && b > 9 && b < 100 {
        	sum += b
        }
    }
    
    fmt.Println(sum)
}