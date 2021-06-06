package main

import "fmt"

func main() {
	var N, num int
	cnt := 0
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&num)
		if num == 0 {
			cnt++
		}
	}
	fmt.Print(cnt)
}
