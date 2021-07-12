package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "unicode"
)

func main() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    textRune := []rune(strings.TrimSpace(text))
    firstLetter := unicode.ToUpper(textRune[0])

    if firstLetter == textRune[0] && textRune[len(textRune) - 1] == '.' {
        fmt.Println("Right")
    } else {
    	fmt.Println("Wrong")
    }
}




