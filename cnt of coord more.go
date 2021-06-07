package main

import "fmt"

func main() {
	var xx, yy, zz, a int
	fmt.Scan(&xx, &yy, &zz)
	a = vote(xx, yy, zz)
	fmt.Print(a)
}

func vote(x int, y int, z int) int {
	var slice []int
	cnt1 := 0
	cnt0 := 0
	slice = append(slice, x, y, z)
	for i := 0; i < 3; i++ {
		if slice[i] == 0 {
			cnt0++
		} else {
			cnt1++
		}
	}
	if cnt0 > cnt1 {
		return 0
	} else {
		return 1
	}
}
