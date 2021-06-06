package main

import "fmt"

func main() {
	var num, num1 int
	fmt.Scan(&num)
	for i := 0; i < 3; i++ {
		num1 = num % 10
		num = num / 10
		fmt.Print(num1)
	}
}
