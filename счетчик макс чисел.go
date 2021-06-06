package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	max := num
	cnt := 0
	for num != 0 {
		if num > max {
			max = num
			cnt = 1
		} else if max == num {
			cnt++
		}
		fmt.Scan(&num)
	}
	fmt.Println(cnt)
}
