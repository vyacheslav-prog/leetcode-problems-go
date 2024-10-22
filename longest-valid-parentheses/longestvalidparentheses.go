package longestvalidparentheses

const (
	closingBracket = ')'
	openingBracket = '('
)

func longestValidParentheses(s string) int {
	brackets := make(map[rune]int)
	for _, value := range s {
		if openingBracket == value {
			brackets[value] += 1
		} else if openingNums := brackets[openingBracket]; 0 != openingNums {
			brackets[value] += 1
		}
	}
	if brackets[closingBracket] < brackets[openingBracket] {
		return 2 * brackets[closingBracket]
	}
	return 2 * brackets[openingBracket]
}
