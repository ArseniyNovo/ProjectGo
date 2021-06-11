package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)
	for i := 1; i < len(text); i++ {
		if i%2 != 0 {
			fmt.Print(string(text[i]))
		}
	}
}
