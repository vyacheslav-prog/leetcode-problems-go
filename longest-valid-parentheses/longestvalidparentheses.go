package longestvalidparentheses

const (
	closingBracket = ')'
	openingBracket = '('
)

func longestValidParentheses(s string) int {
	var leftSubstringIndex, maxLength, rightSubstringIndex, substringLength int
	brackets := make(map[rune]int)
	for index, value := range s {
		if openingBracket == value {
			brackets[value] += 1
		} else if openingNums := brackets[openingBracket]; 0 != openingNums {
			brackets[value] += 1
		}
		if closingNums, openingNums := brackets[closingBracket], brackets[openingBracket]; closingNums < openingNums {
			rightSubstringIndex = index
			if 0 != rightSubstringIndex {
				if closingBracket == value {
					leftSubstringIndex -= 1
				} else {
					leftSubstringIndex += 1
				}
				if closingBracket == s[rightSubstringIndex] {
					substringLength = rightSubstringIndex - leftSubstringIndex
				}
			}
		} else {
			leftSubstringIndex, rightSubstringIndex = index, index
			substringLength = 2 * openingNums
			if openingNums < closingNums {
				brackets = make(map[rune]int)
			}
		}
		if maxLength < substringLength {
			maxLength = substringLength
		}
	}
	return maxLength
}
