package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	b = a % 10
	a = a / 10
	c = a % 10
	a = (a / 10) % 10
	if a != b && a != c && b != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
