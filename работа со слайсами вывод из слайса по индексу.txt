package main

import "fmt"

func main() {
	var numN, num int
	var slice []int
	fmt.Scan(&numN)
	for i := 0; i < numN; i++ {
		fmt.Scan(&num)
		slice = append(slice, num)
	}
	fmt.Print(slice[3])
}
