package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint
	fmt.Scan(&a)

  fn := func(num uint) uint {
    var res string
    strNum := strconv.FormatUint(uint64(num), 10)
    for i := 0; i < len(strNum); i++ {
      numUin64, _ := strconv.ParseUint(string(strNum[i]), 10, 64)
      if numUin64%2 == 0 && numUin64 != 0 {
        res += string(strNum[i])
      }
    }

    numUin64, _ := strconv.ParseUint(res, 10, 64)
    if res == "" {
      numUin64 = 100
    }

    return uint(numUin64)
  }


	fmt.Println(fn(a))
}