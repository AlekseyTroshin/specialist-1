package main

import (
	"fmt"
)

func main() {
	fmt.Println(vote(1, 0, 1))
}

func vote(x int, y int, z int) int {
    a, b, res := 0, 0, 0

    if x == 0 {
         a++   
    } else {
         b++   
    }
    if y == 0 {
         a++   
    } else {
         b++   
    }
    if z == 0 {
         a++   
    } else {
         b++   
    }
    
    if a > b {
        res = 0  
    } else {
         res = 1   
    }
    
    return res
}




