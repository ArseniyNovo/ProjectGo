package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	for i := 1; i <= num; {
		fmt.Print(i, " ")
		i = i * 2
	}
}
