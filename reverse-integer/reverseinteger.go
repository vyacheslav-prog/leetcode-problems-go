package reverseinteger

func reverse(x int) int {
	if 0 == (x / 10) {
		return x
	}
	return (x % 10 * 10) + (x / 10)
}
