package main

import "fmt"

func main() {
	var a int
	if fmt.Scan(&a); a > 0 {
		fmt.Println("Число положительное")
	} else if a < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}
