package romantointeger

var symbolToValueMap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var result, sum int
	for _, symbol := range s {
		value := symbolToValueMap[symbol]
		if sum < value {
			result -= sum * 2
		}
		if value < sum {
			sum = value
		} else {
			sum += value
		}
		result += value
	}
	return result
}
