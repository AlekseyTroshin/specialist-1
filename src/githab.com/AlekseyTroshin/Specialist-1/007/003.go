package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word, secondWord, firstLetter, lastLetter string
	var count int
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		word = input.Text()
		if count == 0 {
			secondWord = word
			count++
			continue
		}

		runeWord := []rune(word)
		runeWordSecond := []rune(secondWord)

		firstLetter = string(runeWord[0])
		lastLetter = string(runeWordSecond[len(runeWordSecond)-1])

		if firstLetter == lastLetter {
			secondWord = word
		} else {
			fmt.Println(word)
			break
		}
	}

}
