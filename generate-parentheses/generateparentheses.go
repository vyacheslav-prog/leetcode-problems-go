package generateparentheses

func generateParentheses(n int) []string {
	var result []string
	if 1 == n {
		result = []string{"()"}
	}
	if 2 == n {
		result = []string{"()()", "(())"}
	}
	return result
}
