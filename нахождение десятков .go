package main

import "fmt"

func main() {

	var a, b int32
	fmt.Scan(&a)
	a = a / 10
	b = a % 10
	fmt.Println(b)
}
