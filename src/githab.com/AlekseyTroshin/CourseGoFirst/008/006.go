package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	var strRune = []rune{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arrRune := []rune(scanner.Text())
	strRune = arrRune[:]
	for i := 0; i < len(arrRune); i++ {
		if i == 0 {
			fmt.Println(string(strRune))
		}
		if i%2 == 0 {
			str = string(strRune[2:])
		} else {
			str = string(strRune[:len(strRune)-2])
		}
		strRune = []rune(str)
		fmt.Println(str)
		if len(strRune) <= 2 {
			break
		}
	}

}
