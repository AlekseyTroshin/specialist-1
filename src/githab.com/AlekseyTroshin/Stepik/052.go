package main

import (
    "errors"
    "fmt"
)

func main() {
  res, err := divide()
  if err != nil {
    fmt.Println(errors.New("ошибка"))
  } else {
    fmt.Println(res)
  }

  
}

func divide () (int, error) {
  var a, b int
  _, err := fmt.Scan(&a)
  _, err = fmt.Scan(&b)

  return a / b, err
}