package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Scan(&text)
	for i := 0; i < len(text); i++ {
		if strings.Count(text, string(text[i])) > 1 {
			text = strings.ReplaceAll(text, string(text[i]), "")
		}
	}
	fmt.Print(text)
}
