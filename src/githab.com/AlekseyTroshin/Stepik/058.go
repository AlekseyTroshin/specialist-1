package main

import (
  "fmt"
)

type batteryForTest struct {
  Charge string
}

func (b batteryForTest) String() string {
  var begin, end string
  begin = "["
  for i := 0; i < 10; i++ {
    if string(b.Charge[i]) == "1" {
      end += "X"
    } else {
      begin += " "
    }
  }
  end += "]"

  return fmt.Sprintf("%s%s", begin, end)
}

func main() {
  var str string
  fmt.Scan(&str)
  
  b := batteryForTest{
    Charge: str,
  }
  fmt.Println(b)
}