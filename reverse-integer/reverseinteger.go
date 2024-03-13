package reverseinteger

const INT_MAX = 2147483647

const INT_MIN = -2147483648

func reverse(x int) int {
	var reversing int
	for remainder := x; remainder != 0; remainder /= 10 {
		if (remainder/10 >= INT_MAX/10 || (remainder/1000000000 < 10 && remainder/1000000000 != 0 && remainder%10 > 2)) ||
			(remainder/10 <= INT_MIN/10 || (remainder/1000000000 > -10 && remainder/1000000000 != 0 && remainder%10 < -2)) {
			return 0
		}
		reversing = (10 * reversing) + (remainder % 10)
	}
	return reversing
}
