package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
  var pass string
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
  	str := scanner.Text()
  	if pass == "" {
  		pass = str
  	} else if len(str) < 8 {
  		fmt.Println("Слишком короткий пароль!")
  		pass = ""
  	} else if strings.Contains(str, "qwe") || strings.Contains(str, "123") {
    	fmt.Println("Слишком простой пароль!")
    	pass = ""
    } else if str != pass {
    	fmt.Println("Введенные пароли различаются!")
		pass = ""
    } else if str == pass {
    	fmt.Println("Ну наконец-то!")
    	return
    }
  }
}