package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	if int(a[0]+a[1]+a[2]) == int(a[3]+a[4]+a[5]) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
