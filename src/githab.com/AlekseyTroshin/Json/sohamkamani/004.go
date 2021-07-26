package main

import (
	"fmt"
	"encoding/json"
)

type Dimensions struct {
	Height int
	Width int
}

type Bird struct {
	Species string
	Description string
	Dimensions Dimensions
}

func main() {
	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("%v\n", bird)
}