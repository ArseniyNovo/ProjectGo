func test(x1 *int, x2 *int) {
	buffer := *x1
	*x1 = *x2
	*x2 = buffer
	fmt.Print(*x1, *x2)
}