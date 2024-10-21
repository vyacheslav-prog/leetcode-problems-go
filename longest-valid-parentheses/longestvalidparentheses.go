package longestvalidparentheses

func longestValidParentheses(s string) int {
	var closingNums, openingNums int
	for _, value := range s {
		if '(' == value {
			openingNums += 1
		} else {
			closingNums += 1
		}
	}
	if openingNums < closingNums {
		return 2 * openingNums
	}
	return 2 * openingNums
}
