func T() float64 {
	m := M(p, v)
	return 6 / W(m)
}
func W(m float64) float64 {
	return math.Sqrt(k / m)
}
func M() float64 {
	return p * v
}