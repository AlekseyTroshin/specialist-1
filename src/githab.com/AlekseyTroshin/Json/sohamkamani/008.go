package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)

	bird := result["birds"].(map[string]interface{})

	for key, j := range bird {
		fmt.Println(key, j.(string))
	}
}