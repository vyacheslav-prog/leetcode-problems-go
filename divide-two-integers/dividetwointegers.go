package dividetwointegers

func divide(dividend int, divisor int) int {
	var result int
	for remainder := dividend; divisor <= remainder; remainder -= divisor {
		result += 1
	}
	return result
}
