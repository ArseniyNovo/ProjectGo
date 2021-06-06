package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	switch {
	case (1 == num || num%10 == 1) && (num > 20 || num < 2):
		fmt.Print(num, " korova")
	case (num%10 == 2 || num%10 == 3 || num%10 == 4) && (num > 20 || num < 5):
		fmt.Print(num, " korovy")
	case num < 21 || num > 4 || num%10 == 0 || num%10 == 5 || num%10 == 6 || num%10 == 7 || num%10 == 8 || num%10 == 9:
		fmt.Print(num, " korov")
	}
}
