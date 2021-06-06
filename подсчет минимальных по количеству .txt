package main

import "fmt"

func main() {
	var N, num int
	cnt := 1
	fmt.Scan(&N, &num)
	min := num
    for i := 1; i < N; i++ {
		fmt.Scan(&num)
		if num < min {
			cnt = 0
            min = num
		}
		if min == num {
			cnt++
		}
	}
	fmt.Print(cnt)
}
