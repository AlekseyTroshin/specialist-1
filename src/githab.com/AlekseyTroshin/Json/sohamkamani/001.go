package main

import (
	"fmt"
	"encoding/json"
)

type MyStoredVariable struct {
	Some string
}

func main() {
	myJsonString := `{"some":"json"}`

	var myStore MyStoredVariable

	json.Unmarshal([]byte(myJsonString), &myStore)

	fmt.Printf("%s %T\n", myStore.Some, myStore)


}