package main

import (
	"fmt"
	"bufio"
	"os"
  "strconv"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  var count int
  const minConst, maxConst, maxConstIn = 100.0, 140.0, 141.0
  var min, max = minConst, maxConst

  for scanner.Scan() {
  	str := scanner.Text()
    num, _ := strconv.ParseFloat(str, 64)

    if count == 0 && num >= minConst && num < maxConstIn {
      min, max = num, num
    }

    if min > num && minConst <= num {
      min = num 
    } else if max < num && maxConstIn > num {
      max = num
    }

    if num < 0 {
      fmt.Printf("%d\n%.1f %.1f\n", count, min, max)
      return
    } else if  num >= minConst && num < maxConstIn {
    	count++
    }
  }
}