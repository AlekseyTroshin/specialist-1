// берём значение из ch1_dup2/one.txt ch1_dup2/two.txt
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileName := map[string]map[string]int{}
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++

		}
		fileName[filename] = counts
		counts = make(map[string]int)
	}
	for fName, fMap := range fileName {
		fmt.Println(fName)
		for line, n := range fMap {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
		fmt.Println("-----")
	}
}
