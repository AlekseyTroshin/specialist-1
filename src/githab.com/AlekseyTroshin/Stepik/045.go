package main

import (
    "fmt"
    "strings"
)

func main() {
    var a, b string
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Println(strings.Index(a, b))
}