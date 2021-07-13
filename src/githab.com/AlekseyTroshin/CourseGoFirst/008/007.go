package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count, max int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arrRune := []rune(scanner.Text())

	for i := 0; i < len(arrRune); i++ {
		if arrRune[i] == 'о' {
			count++
		} else {
			count = 0
		}

		if max < count {
			max = count
		}
	}

	fmt.Println(max)
}
