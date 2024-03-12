package reverseinteger

const INT_MAX = 2147483647

const INT_MIN = -2147483648

func reverse(x int) int {
	decimals := 1
	for withoutRightDigit := x / 10; withoutRightDigit != 0; withoutRightDigit /= 10 {
		decimals *= 10
	}
	if (x%decimals > INT_MAX%decimals) && (x/10 >= INT_MAX/10 || INT_MAX%decimals > 7) {
		return 0
	}
	if (x%decimals < INT_MIN%decimals) && (x/10 <= INT_MIN/10 || INT_MIN%decimals < -8) {
		return 0
	}
	var reversing int
	for remainder := x; decimals != 0; decimals /= 10 {
		reversing += remainder % 10 * decimals
		remainder /= 10
	}
	return reversing
}
