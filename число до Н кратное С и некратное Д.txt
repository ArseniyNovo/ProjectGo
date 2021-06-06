package main

import "fmt"

func main() {
	var num, c, d int
	fmt.Scan(&num, &c, &d)
	for i := 1; i <= num; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}
