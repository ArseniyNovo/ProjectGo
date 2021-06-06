package main

import "fmt"

func main() {
	var num float64
	fmt.Scan(&num)
	if num <= 0 {
		fmt.Printf("число %2.2f не подходит", num)
	} else if num > 0 && num < 10000 {
		fmt.Printf("%.4f", num*num)
	} else {
		fmt.Printf("%e", num)
	}
}
