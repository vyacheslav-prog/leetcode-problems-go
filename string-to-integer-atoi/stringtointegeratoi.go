package stringtointegeratoi

var charToDigitMap = map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == ' ' {
		return myAtoi(s[1:])
	}
	if s[0] == '-' {
		return myAtoi(s[1:]) * -1
	}
	var result int
	for position, awaitDigit := 0, true; awaitDigit && position != len(s); position += 1 {
		value, isDigit := charToDigitMap[s[position]]
		if isDigit {
			result *= 10
			result += value
		}
		awaitDigit = isDigit
	}
	return result
}
