package longestvalidparentheses

const (
	closingBracket = ')'
	openingBracket = '('
)

func longestValidParentheses(s string) int {
	var accClosingNums, accOpeningNums, leftOpeningNums, leftUnclosedIndex, maxLength, substringLength int
	for index, value := range s {
		if openingBracket == value {
			accOpeningNums += 1
		} else if 0 != accOpeningNums {
			accClosingNums += 1
		}
		if accOpeningNums <= accClosingNums {
			leftOpeningNums, leftUnclosedIndex = 0, index
			substringLength = 2 * accOpeningNums
			if accOpeningNums != accClosingNums {
				accClosingNums, accOpeningNums = 0, 0
			}
		} else if 0 != index {
			if closingBracket == value {
				substringLength = index - leftUnclosedIndex + 1
				leftUnclosedIndex -= leftOpeningNums
				leftOpeningNums -= 1
			} else {
				leftUnclosedIndex += 1
				if openingBracket == s[index-1] {
					leftOpeningNums += 1
					leftUnclosedIndex = index
				}
			}
		}
		if maxLength < substringLength {
			maxLength = substringLength
		}
	}
	return maxLength
}
