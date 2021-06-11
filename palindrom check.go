package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	text, _ := (bufio.NewReader(os.Stdin).ReadString('\n'))
	textR := []rune(text)
	lenght := utf8.RuneCountInString(text) - 1
	cnt := lenght/2 + 1
	for i := 0; i <= cnt; i++ {
		if textR[i] == textR[lenght-1] {
			if i == lenght-1 || i-1 == lenght-1 {
				fmt.Print("Палиндром")
			}
			lenght--
			continue
		} else {
			fmt.Print("Нет")
			break
		}
	}
}
