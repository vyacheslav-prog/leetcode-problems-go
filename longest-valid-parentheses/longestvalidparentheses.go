package longestvalidparentheses

const (
	closingBracket = ')'
	openingBracket = '('
)

func longestValidParentheses(s string) int {
	var maxLength, substringLength int
	brackets := make(map[rune]int)
	for _, value := range s {
		if openingBracket == value {
			brackets[value] += 1
		} else if openingNums := brackets[openingBracket]; 0 != openingNums {
			brackets[value] += 1
		}
		if closingNums, openingNums := brackets[closingBracket], brackets[openingBracket]; openingNums < closingNums {
			substringLength = 2 * openingNums
			brackets = make(map[rune]int)
		} else if closingNums == openingNums {
			substringLength = 2 * openingNums
		} else {
			substringLength = 2 * closingNums
		}
		if maxLength < substringLength {
			maxLength = substringLength
		}
	}
	return maxLength
}
