package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word, firstWord, secondWord, firstLetter, lastLetter, softSign string
	var count int
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		word = input.Text()
		firstWord = word
		runeWord := []rune(firstWord)
		softSign = string(runeWord[len(runeWord)-1])

		if softSign == "ь" {
			firstWord = string(runeWord[:len(runeWord)-1])
		}

		if count == 0 {
			secondWord = firstWord
			count++
			continue
		}

		runeWordSecond := []rune(secondWord)

		firstLetter = string(runeWord[0])
		lastLetter = string(runeWordSecond[len(runeWordSecond)-1])

		if firstLetter == lastLetter {
			secondWord = firstWord
		} else {
			fmt.Println(word)
			break
		}
	}

}