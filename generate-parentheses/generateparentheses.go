package generateparentheses

func generateParenthesesTail(closedLeft int, openedLeft int, tail string) []string {
	var result []string
	if 0 != openedLeft {
		result = append(result, generateParenthesesTail(closedLeft, openedLeft-1, tail+"(")...)
	}
	if openedLeft < closedLeft {
		result = append(result, generateParenthesesTail(closedLeft-1, openedLeft, tail+")")...)
	}
	if 0 == len(result) {
		result = []string{tail}
	}
	return result
}

func generateParenthesis(n int) []string {
	var result []string
	if 0 < n {
		result = generateParenthesesTail(n, n, "")
	}
	return result
}
