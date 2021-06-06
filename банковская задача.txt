package main

import "fmt"

func main() {
	var cost, perc, sum float32
	fmt.Scan(&cost, &perc, &sum)
	cost = cost * (1.0 + perc/100)
	cnt := 1
	for cost <= sum {
		cnt++
		cost = cost * (1.0 + perc/100)
	}
	fmt.Println(cnt)
}
