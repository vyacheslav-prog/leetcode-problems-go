package longestvalidparentheses

const (
	closingBracket = ')'
	openingBracket = '('
)

func longestValidParentheses(s string) int {
	var accClosingNums, accOpeningNums, leftUnclosedIndex, maxLength, substringLength int
	for index, value := range s {
		if openingBracket == value {
			accOpeningNums += 1
		} else if 0 != accOpeningNums {
			accClosingNums += 1
		}
		if accOpeningNums <= accClosingNums {
			leftUnclosedIndex = index
			substringLength = 2 * accOpeningNums
			if accOpeningNums != accClosingNums {
				accClosingNums, accOpeningNums = 0, 0
			}
		} else {
			if leftUnclosedIndex != index {
				if closingBracket == value {
					leftUnclosedIndex -= 1
				} else {
					leftUnclosedIndex += 1
				}
				if closingBracket == value {
					substringLength = index - leftUnclosedIndex
				}
			}
		}
		if maxLength < substringLength {
			maxLength = substringLength
		}
	}
	return maxLength
}
