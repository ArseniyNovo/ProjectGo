package main

import "fmt"

func main() {
	var num, sum int
	sum = 0
	fmt.Scan(&num)
	for i := 0; i < 3; i++ {
		sum = sum + num%10
		num = num / 10
	}
	fmt.Print(sum)
}
