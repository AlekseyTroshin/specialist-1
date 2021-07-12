package main

import (
	"fmt"
)

func main() {
	r := '2'
	num := 1023
	strNum := "2222"
	str := "hello"

	fmt.Println("1", string(r))
	fmt.Println("2", int(r))
	fmt.Println("3", r)
	fmt.Println("4", num)
	fmt.Println("5", strNum[0])
	fmt.Println("6", string(strNum[0]))
	fmt.Println("6.1", int(strNum[0]))
	fmt.Println("7", str[0])
	fmt.Println("8", string(str[0]))
}
