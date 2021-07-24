package main

import (
	"fmt"
	"encoding/json"
)

type Bird struct {
	Species string
	Description string
}

func main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
}