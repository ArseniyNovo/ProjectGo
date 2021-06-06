package main

import "fmt"

func main() {
	var num1, num2 float32
	fmt.Scan(&num1, &num2)
	num1 = (num1 + num2) / 2
	fmt.Print(num1)

}
