package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

type Average struct {
	Average float32
}

func main() {
	file, err := os.Open("./json/students.json")
	data, err := ioutil.ReadAll(file)

	var countStudent, sumRating float64
	
	if err != nil {
	    panic(err)
	}

	var jsonA map[string]interface{}

	if err := json.Unmarshal(data, &jsonA); err != nil {
		fmt.Println(err)
	}
	
	a := jsonA["Students"].([]interface{})
	for _, i := range a {
		z := i.(map[string]interface{})
		zz := z["Rating"].([]interface{})
		for range zz {
			sumRating++
		}
		countStudent++
	}

	gg := Average {
		Average: float32(sumRating / countStudent),
	}

	b, err := json.MarshalIndent(gg, "", "    ")
	
	fmt.Println(string(b))
}
