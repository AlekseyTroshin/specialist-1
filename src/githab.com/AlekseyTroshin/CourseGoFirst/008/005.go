package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arrRune := []rune(scanner.Text())

	for i := 0; i < len(arrRune); i++ {
		if i%2 == 0 {
			str += string(arrRune[i]) + string(arrRune[i]) + string(arrRune[i])
		}
	}
	fmt.Println(str)
}
