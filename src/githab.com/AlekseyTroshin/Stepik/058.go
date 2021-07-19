package main

import "fmt" 

type Battery struct {
  Charge string
}

func (b Battery) String() string {
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
  
  batteryForTest := Battery{
    Charge: str,
  }