// задача H
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var a, b, c int
	scanner.Scan()
	a, _ = strconv.Atoi(scanner.Text())
	listWork := []string{}
	listDo := []int{}

	for i := 0; i < a; i++ {
		scanner.Scan()
		listWork = append(listWork, scanner.Text())
	}

	scanner.Scan()
	b, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < b; i++ {
		scanner.Scan()
		c, _ = strconv.Atoi(scanner.Text())
		listDo = append(listDo, c-1)
	}

	for _, item := range listDo {
		fmt.Println(listWork[item])
	}
}
