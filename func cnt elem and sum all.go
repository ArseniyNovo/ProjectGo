package main

func main() {

}

func sumInt(n ...int) (cnt, sum int) {
	sum = 0
	cnt = 0
	for _, elem := range n {
		cnt++
		sum = sum + elem
	}
}
