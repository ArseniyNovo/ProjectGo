package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	res, err := divide(num1, num2)
	if err != nil {
		fmt.Print("ошибка")
	} else {
		fmt.Print(res)
	}
}
