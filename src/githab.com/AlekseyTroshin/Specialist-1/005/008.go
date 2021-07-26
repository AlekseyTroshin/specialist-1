package main

import (
    "fmt"
)

func main() {
    var a, b, c int
    fmt.Scan(&a)
    for i := 1; i <= a; i++ {
        fmt.Scan(&b)
        if i % 2 == 0 {
            c += -b
        } else {
            c += b
        }
    }
    fmt.Println(c)
}