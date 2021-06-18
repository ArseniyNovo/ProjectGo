package main

import (
	"fmt"
)

func main() {
	var num int
	var arrayNum []int
	fmt.Scan(&num)
	for num/10 != 0 || num%10 != 0 {
		arrayNum = append(arrayNum, (num%10)*(num%10))
		num /= 10
	}
	for i := len(arrayNum) - 1; i >= 0; i-- {
		fmt.Print(arrayNum[i])
	}
}
