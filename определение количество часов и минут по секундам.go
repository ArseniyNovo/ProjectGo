package main

import "fmt"

func main() {
	var sec int32
	fmt.Scan(&sec)
	hour := sec / 3600
	sec = sec % 3600
	min := sec / 60
	fmt.Printf("It is %d hours %d minutes.", hour, min)
}