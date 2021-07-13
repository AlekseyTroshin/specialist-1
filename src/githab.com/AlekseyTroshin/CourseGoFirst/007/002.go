package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str, beginStr, endStr string
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str = input.Text()
	strRune := []rune(str)
	beginStr = string(strRune[0])
	endStr = string(strRune[len(strRune)-1])
	beginStr = strings.ToLower(beginStr)
	endStr = strings.ToLower(endStr)

	if beginStr == "д" && endStr == "а" || beginStr == "а" && endStr == "д" {
		fmt.Println("СОГЛАСЕН")
	} else {
		fmt.Println("НЕ СОГЛАСЕН")
	}

}
