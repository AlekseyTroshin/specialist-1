package main

import (
  "fmt"
)

func main() {
  var i int
  _, err := fmt.Scan(&i)
  if err != nil {
    fmt.Println("no no no no", i)
  } else {
    fmt.Println("yes yes yes", i)
  }
}