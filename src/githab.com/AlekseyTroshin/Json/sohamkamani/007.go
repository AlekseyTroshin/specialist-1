package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Global struct {
	GlobalId float64 `json:"global_id"`
}

func main() {
	file, err := os.Open("../datatest/data-20190514T0100.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	var global []Global
	json.Unmarshal(bytes, &global)
	fmt.Printf("%T", global)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", global[i])
	}

}
