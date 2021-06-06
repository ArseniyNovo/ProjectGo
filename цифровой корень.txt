package main

import "fmt"

func main() {
	var num, sum int64
	sum = 0
	fmt.Scan(&num)
	for {
		if num/10 != 0 {
			sum = sum + num%10
			num = num / 10
			continue
		}
		sum = sum + num
		if sum/10 != 0 {
			sum = sum/10 + sum%10
			fmt.Print(sum)
			break
		} else {
			fmt.Print(sum)
			break
		}
	}
}
