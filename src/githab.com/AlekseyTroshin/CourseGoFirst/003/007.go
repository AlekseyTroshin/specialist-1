package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fmt.Println(strings.Contains(sc.Text(), "a"))
}