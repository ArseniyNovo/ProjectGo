package main

import (
	"fmt"
	"time"
)

func work(x int) int {
	time.Sleep(time.Second)
	if x >= 4 {
		return x + 1
	} else {
		return x - 1
	}
}
func main() {
	var num int
	number := make(map[int]int)
	for i := 0; i < 10; i++ {
		fmt.Scan(&num)
		if buff, inMap := number[num]; inMap {
			fmt.Print(buff, " ")
		} else {
			buff := work(num)
			number[num] = buff
			fmt.Print(buff, " ")
		}
	}
}
