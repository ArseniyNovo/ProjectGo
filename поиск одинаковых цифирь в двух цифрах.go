package main

import "fmt"

func main() {
	var num1, num2 string
	fmt.Scan(&num1, &num2)
	for i := 0; i < len(num1); {
		for j := 0; j < len(num2); {
			if num1[i] == num2[j] {
				fmt.Print(string(num1[i]), " ")
			}
			j++
		}
		i++
	}
}
