package longestvalidparentheses

const (
	closingBracket = ')'
	openingBracket = '('
)

func longestValidParentheses(s string) int {
	var closingCounter, maxLength, openingCounter int
	for i := 0; i < len(s); i += 1 {
		if openingBracket == s[i] {
			openingCounter += 1
		} else {
			closingCounter += 1
		}
		if openingCounter == closingCounter {
			if currentLength := 2 * closingCounter; maxLength < currentLength {
				maxLength = currentLength
			}
		} else if openingCounter < closingCounter {
			openingCounter, closingCounter = 0, 0
		}
	}
	openingCounter, closingCounter = 0, 0
	for i := len(s) - 1; 0 <= i; i -= 1 {
		if openingBracket == s[i] {
			openingCounter += 1
		} else {
			closingCounter += 1
		}
		if openingCounter == closingCounter {
			if currentLength := 2 * openingCounter; maxLength < currentLength {
				maxLength = currentLength
			}
		} else if closingCounter < openingCounter {
			openingCounter, closingCounter = 0, 0
		}
	}
	return maxLength
}
