package longestvalidparentheses

const (
	closingBracket = ')'
	openingBracket = '('
)

func findLongestValidParentheses(prevResults map[string]int, s string) int {
	if prevResult, isFound := prevResults[s]; isFound {
		return prevResult
	}
	var maxLength, substringLength int
	brackets := make(map[rune]int)
	for index, value := range s {
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
			substringLength = findLongestValidParentheses(prevResults, s[index+1:])
		}
		if maxLength < substringLength {
			maxLength = substringLength
		}
	}
	prevResults[s] = maxLength
	return maxLength
}

func longestValidParentheses(s string) int {
	prevResults := make(map[string]int)
	return findLongestValidParentheses(prevResults, s)
}
