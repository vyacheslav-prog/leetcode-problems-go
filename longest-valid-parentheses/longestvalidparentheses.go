package longestvalidparentheses

const (
	closingBracket = ')'
	openingBracket = '('
)

func longestValidParentheses(s string) int {
	var maxLength, openedSubstringIndex, openedSubstringCounter, substringLength int
	brackets := make(map[rune]int)
	for index, value := range s {
		if openingBracket == value {
			brackets[value] += 1
		} else if openingNums := brackets[openingBracket]; 0 != openingNums {
			brackets[value] += 1
		}
		if closingNums, openingNums := brackets[closingBracket], brackets[openingBracket]; closingNums < openingNums {
			if openingBracket == value {
				openedSubstringCounter += 1
			} else {
				openedSubstringCounter -= 1
			}
			if 1 < openedSubstringCounter {
				openedSubstringCounter = 1
				openedSubstringIndex = index
			} else if 0 == openedSubstringCounter {
				substringLength = index - openedSubstringIndex + 1
			}
		} else {
			openedSubstringCounter = 0
			substringLength = 2 * openingNums
			if openingNums < closingNums {
				brackets = make(map[rune]int)
			}
		}
		if maxLength < substringLength {
			maxLength = substringLength
		}
	}
	if openedSubstringLength := len(s) - openedSubstringIndex + 1; 0 != openedSubstringIndex && 0 != openedSubstringCounter && maxLength < openedSubstringLength {
		maxLength = openedSubstringLength
	}
	return maxLength
}
