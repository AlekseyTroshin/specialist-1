package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

type Global struct {
	GlobalId float64 `json:"global_id"`
}

func main() {
	file, err := os.Open("../datatest/data-20190514T0100.json")
	bytes, err := ioutil.ReadAll(file)
		if err != nil {
	    panic(err)
	}
	var global []Global
	json.Unmarshal(bytes, &global)
	fmt.Printf("%T", global)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", global[i])
	}
	
}