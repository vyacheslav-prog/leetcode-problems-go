package stringtointegeratoi

var charToDigitMap = map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

const (
	maxNegativeNum = 2147483648
	maxPositiveNum = 2147483647
	negativeMul    = -1
	positiveMul    = 1
)

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == ' ' {
		return myAtoi(s[1:])
	}
	maxNumber, mul := maxPositiveNum, positiveMul
	var startPosition, result int
	if s[0] == '-' {
		maxNumber = maxNegativeNum
		mul = negativeMul
		startPosition = 1
	} else if s[0] == '+' {
		startPosition = 1
	}
	for awaitDigit, position := true, startPosition; awaitDigit && position != len(s); position += 1 {
		value, isDigit := charToDigitMap[s[position]]
		if isDigit {
			result *= 10
			result += value
		}
		awaitDigit = isDigit && result < maxNumber
	}
	return min(maxNumber, result) * mul
}
