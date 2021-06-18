package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string
	fmt.Scan(&str)
	for i := 0; i < utf8.RuneCountInString(str); i++ {
		if i == utf8.RuneCountInString(str)-1 {
			fmt.Print(string(str[i]))
			break
		}
		fmt.Print(string(str[i]), "*")
	}

}
