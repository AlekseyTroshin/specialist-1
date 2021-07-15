package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var textReverse string
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.ToLower(strings.TrimSpace(text))
	textSplit := strings.Split(text, "")

	for i := len(textSplit) - 1; i >= 0; i-- {
		textReverse += textSplit[i]
	}

	if text == textReverse {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
