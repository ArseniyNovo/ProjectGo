package main

import (
	"fmt"
	"strings"
)

func main() {
	var text1, text2 string
	fmt.Scan(&text1, &text2)
	fmt.Print(strings.Index(text1, text2))
}
