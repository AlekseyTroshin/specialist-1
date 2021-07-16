package main

import (
  // "encoding/json" // пакет используется для проверки ответа, не удаляйте его
  "fmt"           // пакет используется для проверки ответа, не удаляйте его
  // "os"            // пакет используется для проверки ответа, не удаляйте его
)



func main() {
    // value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
                                            // все полученные значения имеют тип пустого интерфейса
    var value1 interface{} = 33.33                                   
    var value2 interface{} = 4.123                                    
    var operation interface{} = 88                               
    vA, okA := value1.(float64)
    vB, okB := value2.(float64)
    if !okA {
        checkError(value1)
        return
    }

    if !okB {
        checkError(value2)
        return
    }
    
    if v, ok := operation.(string); ok {
        switch {
            case v == "+":
            fmt.Printf("%.4f\n", vA + vB)
            case v == "-":
            fmt.Printf("%.4f\n", vA - vB)
            case v == "*":
            fmt.Printf("%.4f\n", vA * vB)
            case v == "/":
            fmt.Printf("%.4f\n", vA / vB)
            default:
            fmt.Println("неизвестная операция")
        }
    } else {
        fmt.Println("неизвестная операция")
    }
}
                   
func checkError(value interface{}) {
  fmt.Printf("value=%v: %T\n", value, value)
}