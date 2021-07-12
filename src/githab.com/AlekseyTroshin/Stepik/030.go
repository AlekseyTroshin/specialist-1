package main

import (
	"fmt"
)


func main() {
	var a, b, c, count int
	fmt.Scan(&a)
	
    for i := 0; i < a; i++ {
    	fmt.Scan(&b)
    	if i == 0 {
    		c = b
    	}
        
        if b == c {
            count++
        } else if b < c {
        	c = b
            count = 1   
        } 
    }
    
    fmt.Println(count)
}