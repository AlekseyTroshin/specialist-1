package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var a, b, c int
	fmt.Scan(&a)
	listWork := []string{}
	listDo := []int{}

	for i := 0; i < a; i++ {
		scanner.Scan()
		listWork = append(listWork, scanner.Text())
	}

	fmt.Scan(&b)

	for i := 0; i < b; i++ {
		fmt.Scan(&c)
		listDo = append(listDo, c)
	}

	for _, item := range listDo {
		fmt.Println(listWork[item])
	}
}
