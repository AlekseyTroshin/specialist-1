package main

import "fmt"

func main() {
	var word1, word2, word3, word4, word5 string
    fmt.Scan(&word1, &word2, &word3, &word4, &word5)
    
    fmt.Printf("%s - %s\n", word4, word5)
    fmt.Printf("%s - %s\n", word3, word5)
    fmt.Printf("%s - %s\n", word2, word5)
    fmt.Printf("%s - %s\n", word1, word5)
}