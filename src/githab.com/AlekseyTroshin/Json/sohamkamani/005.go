package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	numberJson := "3"
	floatJson := "3.12345"
	stringJson := `"Hello i'm String"`

	var n int
	var f float64
	var s string

	json.Unmarshal([]byte(numberJson), &n)
	fmt.Println(n)

	json.Unmarshal([]byte(floatJson), &f)
	fmt.Println(f)

	json.Unmarshal([]byte(stringJson), &s)
	fmt.Println(s)
}