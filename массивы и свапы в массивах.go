package main

import "fmt"

func main() {
	var workArray [10]uint8
	var a, b, c uint8
	for i := range workArray {
		fmt.Scan(&workArray[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&a, &b)
		c = workArray[a]
		workArray[a] = workArray[b]
		workArray[b] = c
	}
	for i := range workArray {
		fmt.Print(workArray[i], " ")
	}
}
