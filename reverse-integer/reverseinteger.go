package reverseinteger

const INT_MAX = 2147483647

const INT_MIN = -2147483648

func isNegativeOverflow(remainder, reversing int) bool {
	if reversing < INT_MIN/10 || (reversing == INT_MIN/10 && remainder%10 < -8) {
		return true
	}
	return false
}

func isPositiveOverflow(remainder, reversing int) bool {
	if reversing > INT_MAX/10 || (reversing == INT_MAX/10 && remainder%10 > 7) {
		return true
	}
	return false
}

func reverse(x int) int {
	var reversing int
	isOverflow := isPositiveOverflow
	if x < 0 {
		isOverflow = isNegativeOverflow
	}
	for remainder := x; remainder != 0; remainder /= 10 {
		if isOverflow(remainder, reversing) {
			return 0
		}
		reversing = (10 * reversing) + (remainder % 10)
	}
	return reversing
}
