package main

import "fmt"

func main() {
	var a int
	a = minimumFromFour()
	fmt.Print(a)
}

func minimumFromFour() int {
	var a, min int
	for i := 0; i < 4; i++ {
		fmt.Scan(&a)
		if i == 0 {
			min = a
		}
		if min > a {
			min = a
		}
	}
	return min
}
