package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
  	str := scanner.Text()
    if str == "" {
    	return
    }
    fmt.Println(str)
  }
}