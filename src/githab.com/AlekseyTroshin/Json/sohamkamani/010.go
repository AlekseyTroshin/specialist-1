package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species string `json:"birdType"`
	//hello
	Description string `json:"what it does,omitempty"`
}

func main() {
	pigeon := &Bird {
		Species: "Pigeon",
	}

	data, _ := json.Marshal(pigeon)

	fmt.Println(string(data))
}