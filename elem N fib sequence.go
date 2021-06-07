package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n = fibonacci(n)
	fmt.Print(n)
}

func fibonacci(n int) int {
	var sliceFib []int
	sliceFib = append(sliceFib, 1, 1)
	for i := 2; i < n; i++ {
		sumFib := sliceFib[i-1] + sliceFib[i-2]
		sliceFib = append(sliceFib, sumFib)
	}
	return sliceFib[n-1]
}
