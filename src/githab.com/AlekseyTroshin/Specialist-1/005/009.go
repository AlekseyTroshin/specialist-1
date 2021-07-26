package main

import (
    "fmt"
)

func main() {
    var a, b, xa, ya, xb, yb, xc, yc, xd, yd, count, x, y, max, min int

    for i := 0; i < 4; i++ {
        fmt.Scan(&a, &b)
        if a > b {
            max = a
            min = b
        } else {
            max = b
            min = a
        }

        switch {
        case 1 == i:  
            xa = min
            ya = max
        case 2 == i: 
            xb = min
            yb = max
        case 3 == i: 
            xc = min
            yc = max
        case 4 == i: 
            xd = min
            yd = max
        }
    }
    fmt.Scan(&count)
    for i := 0; i < count; i++ {
        fmt.Scan(&x, &y)
        if (xa <= x && ya >= y) || (xb <= x && yb >= y) || (xc <= x && yc >= y) || (xd <= x && yd >= y) {
            fmt.Printf("Точка с координатами %d,%d принадлежит исследуемому квадрату\n", x, y)
        } else {
            fmt.Printf("Точка с координатами %d,%d не принадлежит исследуемому квадрату\n", x, y)
        }
    }
}