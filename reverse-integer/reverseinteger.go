package reverseinteger

func reverse(x int) int {
	significantDecimals := 1
	for withoutRightDigit := x / 10; withoutRightDigit != 0; withoutRightDigit /= 10 {
		significantDecimals *= 10
	}
	if significantDecimals != 1 {
		return (x % 10 * significantDecimals) + reverse(x/10)
	}
	return x
}
