package main

import "fmt"

func main() {
    var a, b, count int
    b = 0
    count = 0
    
    for {
    	fmt.Scan(&a)
    	if a == 0 {
    		break
    	}
    	if a >= b {
    		if a != b {
    			count = 0
    		}
    		b = a
    		count++
    	} 
    }
    
    fmt.Println(count)
}