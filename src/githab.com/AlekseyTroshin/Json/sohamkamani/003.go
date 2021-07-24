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
	birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("%v\n", birds)
}