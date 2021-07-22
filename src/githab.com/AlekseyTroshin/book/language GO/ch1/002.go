package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args
	for i, val := range s {
		fmt.Println(i+1, val)
	}
}
