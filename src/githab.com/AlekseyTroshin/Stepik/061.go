package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

type (
	Student struct {
		Rating []int
	}
	Group struct {
		Students []Student
	}
)

func main() {
	file, _ := os.Open("./json/students.json")
	bytes, _ := ioutil.ReadAll(file)
	var countStudent int
	var students = new(Group)
	
	if err := json.Unmarshal(bytes, students); err != nil {
		fmt.Println(err)
	}

	for _, student := range students.Students {
		countStudent += len(student.Rating)
	}

	average := float64(countStudent) / float64(len(students.Students))

	bytes, _ = json.MarshalIndent(map[string]float64{"Average": average}, "", "    ")

	fmt.Printf("%s\n", bytes)
}
