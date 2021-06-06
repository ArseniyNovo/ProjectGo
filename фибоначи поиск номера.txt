package main

import "fmt"

func main() {
	var num, fib1, fib2 int
	fmt.Scan(&num)
	fib1 = 1
	fib2 = 1
	if num == 1 {
		fmt.Print("1")
	}
	for i := 3; num >= fib1+fib2; {
		if num == fib1+fib2 {
			fmt.Print(i)
			break
		}
		a := fib1
		fib1 = fib2
		fib2 = a + fib1
		i++
	}
	if num < fib1+fib2{
		fmt.Print("-1")
	}

}
