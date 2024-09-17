package dividetwointegers

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
	for quotient := 31; -1 != quotient; quotient -= 1 {
		if divisor<<quotient <= dividend {
			dividend -= divisor << quotient
			result += 1 << quotient
		}
	}
	if isNegativeDividend != isNegativeDivisor {
		if 0 != result>>31 {
			result = -(1 << 31)
		} else {
			result = -result
		}
	}
	if 0 < result && 0 != result>>31 {
		result = 1<<31 - 1
	}
	return result
}
