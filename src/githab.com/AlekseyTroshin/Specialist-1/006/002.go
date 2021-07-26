// задача A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var a, b int
	fmt.Scan(&a, &b)
	tvSeriesServise := []string{}
	tvSeriesYouFind := []string{}

	for i := 0; i < a; i++ {
		scanner.Scan()
		tvSeriesServise = append(tvSeriesServise, scanner.Text())
	}

	for i := 0; i < b; i++ {
		scanner.Scan()
		tvSeriesYouFind = append(tvSeriesYouFind, scanner.Text())
	}

outerSeriesYouFind:
	for _, seriesYouFind := range tvSeriesYouFind {
		for _, seriesServise := range tvSeriesServise {
			if seriesServise == seriesYouFind {
				fmt.Println("ЕСТЬ")
				continue outerSeriesYouFind
			}
		}
		fmt.Println("НЕТ В НАЛИЧИИ")
	}
}
