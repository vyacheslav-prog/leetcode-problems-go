package dividetwointegers

const (
	negativeMax = -(1 << 31)
	positiveMax = 1 << 31 - 1
)

func divide(dividend int, divisor int) int {
	var isNegativeDividend, isNegativeDivisor bool
	var result int
	if dividend < 0 {
		dividend = -dividend
		isNegativeDividend = true
	}
	if divisor < 0 {
		divisor = -divisor
		isNegativeDivisor = true
	}
	for remainder := dividend; divisor <= remainder; remainder -= divisor {
		result += 1
	}
	if isNegativeDividend != isNegativeDivisor {
		if 0 != result >> 31 {
			result = negativeMax
		} else {
			result = -result
		}
	}
	if 0 < result && 0 != result >> 31 {
		result = positiveMax
	}
	return result
}
