package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	for num2 != num1 {
		if num2%7 == 0 {
			fmt.Print(num2)
			break
		}
		num2--
	}
	if num1 == num2 {
		fmt.Print("NO")
	}
}
