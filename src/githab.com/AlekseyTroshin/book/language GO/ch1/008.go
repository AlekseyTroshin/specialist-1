package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	randomPallette()
}


func randomPallette() {
	var num = []int{}
	for i := 0; i < 4; i++ {
		num = append(num, rand.Intn(255))
	}
	aaa(num)
}

func aaa(a []int) {
	var z = []int{}
	z = append(z, a...)
	fmt.Println(a)
}