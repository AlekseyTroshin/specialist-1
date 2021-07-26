package main

import (
    "fmt"
)

func main() {
    var xa, xb, ya, yb, sum int
    fmt.Scan(&xa, &xb, &ya, &yb)

    xa = abc(xa - ya)
    xb = abc(xb - yb)
    sum = xa + xb

    if (sum > 3 || sum < 3) || xa == 3 || xb == 3 {
        fmt.Println("НЕТ")
    } else {
        fmt.Println("ДА")
    }
}

func abc(a int) int {
    if a < 0 {
        return -a
    }
    return a
}