package main

import "fmt"

func main() {
	var N, num, cnt int
	var slice []int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&num)
		slice = append(slice, num)
		if num > 0 {
			cnt = cnt + 1
		}
	}
	fmt.Print(cnt)
}
