package validparentheses

var bracketToDealMap = map[byte]int{
	'(': 1,
	')': -1,
	'[': 2,
	']': -2,
	'{': 3,
	'}': -3,
}

func isValid(s string) bool {
	var balance int
	for index := 0; len(s) != index; index += 1 {
		deal := bracketToDealMap[s[index]]
		if 0 == balance || 0 == balance+deal {
			balance += deal
		}
	}
	return 0 == balance
}
