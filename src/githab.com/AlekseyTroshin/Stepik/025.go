package main

import "fmt"

func main() {
    var a, b, count int

    fmt.Scan(&a)

    for i := 0; i < a; i++ {
        fmt.Scan(&b)
        if b > 0 {
             count++   
        }
    }
    
    fmt.Println(count)
}