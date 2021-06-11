package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	text, _ := (bufio.NewReader(os.Stdin).ReadString('\n'))
	textR := []rune(text)
	textS := []rune(".")
	fmt.Println(textR)
	fmt.Println(textS)
	if unicode.IsUpper(textR[0]) && textR[utf8.RuneCountInString(text)-2] == textS[0] {
		fmt.Print("Right", "\t", unicode.IsUpper(textR[0]), textR[utf8.RuneCountInString(text)-2] == textS[0], textR[utf8.RuneCountInString(text)-2], textS[0], utf8.RuneCountInString(text))
	} else {
		fmt.Print("Wrong", "\t", unicode.IsUpper(textR[0]), textR[utf8.RuneCountInString(text)-2] == textS[0], textR[utf8.RuneCountInString(text)-2], textS[0], utf8.RuneCountInString(text))
	}
}
