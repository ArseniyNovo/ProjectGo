package main

import "fmt"

func main() {
	var N, num int
	var slice []int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&num)
		slice = append(slice, num)
		if i%2 == 0 {
			fmt.Print(slice[i], " ")
		}
	}
}
