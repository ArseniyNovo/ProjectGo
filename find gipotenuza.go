package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	fmt.Scan(&num1, &num2)
	fmt.Print(math.Sqrt(math.Pow(num1, 2) + math.Pow(num2, 2)))
}
