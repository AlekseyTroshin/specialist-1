package main

import (
    "fmt"
)

func main() {
    var z, a, b, one, two uint8
    workArray := [10]uint8{}
    
    for i := 0; i < len(workArray); i++ {
        fmt.Scan(&z)
        workArray[i] = z
    }
    for i := 0; i < 3; i++ {
        fmt.Scan(&a)
        fmt.Scan(&b)
        one = workArray[a]
        two = workArray[b]
        workArray[a] = two
        workArray[b] = one
    }
    for i := range workArray {
        fmt.Print(workArray[i], " ")
    }
}