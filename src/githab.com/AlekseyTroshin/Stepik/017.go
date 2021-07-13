package main

import (
    "fmt"
    "reflect"
)

func main() {
    a := 4.234522
    b := fmt.Sprintf("%.2f", a)
    fmt.Println(b)
    fmt.Println(reflect.TypeOf(b))
}