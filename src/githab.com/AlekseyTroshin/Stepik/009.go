package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a%400 == 0 {
		fmt.Println("YES")
	} else if a%4 == 0 && a%100 == 0 {
		fmt.Println("NO")
	} else if a%4 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
