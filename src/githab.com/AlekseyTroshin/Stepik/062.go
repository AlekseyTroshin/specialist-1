package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

type (
	GlobalId struct {
		GlobalId int `json:"global_id"`
	}
	Group struct {
		Globals []map[string]GlobalId
	}
)

func main() {
	file, _ := os.Open("./json/data-20190514T0100.json")
	bytes, _ := ioutil.ReadAll(file)

	// var globalid []map[string]interface{}

	var globalid = new(Group)

	if err := json.Unmarshal(bytes, globalid); err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("%v\n", globalid)
}
