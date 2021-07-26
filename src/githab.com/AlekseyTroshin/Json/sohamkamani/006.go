package main

import (
	"fmt"
	"encoding/json"
)

type Bird struct {
	Species string `json:"birdType"`
	Description string `json:"what it does"`
}

func main() {
	birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
}