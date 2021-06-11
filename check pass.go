package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var text string
	fmt.Scan(&text)
	symb := true
	digit := true
	cntDigit := 0
	cntSymb := 0
	textR := []rune(text)
	if len(textR) >= 5 {
		for i := 0; symb && digit && i <= utf8.RuneCountInString(text); i++ {
			if unicode.Is(unicode.Latin, textR[i]) || unicode.IsDigit(textR[i]) {
				if unicode.Is(unicode.Latin, textR[i]) {
					symb = false
					cntSymb++
				}
				if unicode.IsDigit(textR[i]) {
					digit = false
					cntDigit++
				}
			} else {
				fmt.Print("Wrong password")
				break
			}
			if i == utf8.RuneCountInString(text)-1 {
				if (symb == false || cntSymb > 0) && (digit == false || cntDigit > 0) {
					fmt.Print("Ok")
				}
			} else {
				symb = true
				digit = true
			}
		}
	} else {
		fmt.Print("Wrong password1")
	}
}
