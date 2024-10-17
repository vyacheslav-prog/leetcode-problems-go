package longestvalidparentheses

import "testing"

func TestFindsZeroLengthSubstringForEmptyString(t *testing.T) {
	result := longestValidParentheses("")
	if 0 != result {
		t.Errorf("Result must be zero for empty string, actual is [%v]", result)
	}
}
