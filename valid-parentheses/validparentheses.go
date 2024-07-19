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
	var balance, prevDeal int
	for index := 0; len(s) != index; index += 1 {
		deal := bracketToDealMap[s[index]]
		if 0 <= deal+prevDeal {
			balance += deal
		}
		prevDeal = deal
	}
	return 0 == balance
}
