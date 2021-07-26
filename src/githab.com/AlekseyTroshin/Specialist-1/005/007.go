package main

import (
    "fmt"
    "math"
)

func main() {
    var a float64
    fmt.Scan(&a)
    for i := -a; i <= a; i++ {
        fmt.Printf("Квадрат числа %d равен %d\n", int(i), int(math.Pow(i, 2)))
    }
}