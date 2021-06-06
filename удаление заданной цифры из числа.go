package main

import "fmt"

func main() {
	var num, numd, a int
	fmt.Scan(&num, &numd)
	a = 1
	num1 := 0
	for num/10 != 0 {
		if num%10 != numd {
			num1 = (num%10)*a + num1
			a = a * 10
		}
		num /= 10
	}
	if num%10 != numd {
		num1 = num1 + (num%10)*a
	}
	fmt.Print(num1)
}
