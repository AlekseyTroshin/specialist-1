package main

import (
    "fmt"
)

func main() {
    var a, count int
    fmt.Scan(&a)
    if a == 1 {
        fmt.Println(a)
        return
    }
    for i := 1; i <= a; i++ {
        if a % i == 0 {
            if a == i && count < 2 {
                fmt.Printf("%d\nACHTUNG\n", a)
                return
            }
            if i == a {
                fmt.Println(i)
                return
            }
            fmt.Printf("%d ", i)
            count++
        }
    }
}