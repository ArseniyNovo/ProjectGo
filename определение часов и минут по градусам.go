package main

import "fmt"

func main() {
	var d, h, m int
	fmt.Scan(&d)
	h = d / 30
	m = (d - h*30) * 2
	fmt.Println("It is", h, "hours", m, "minutes.")
}