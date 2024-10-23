package longestvalidparentheses

import "testing"

func TestFindsZeroLengthSubstringForEmptyString(t *testing.T) {
	result := longestValidParentheses("")
	if 0 != result {
		t.Errorf("Result must be zero for empty string, actual is [%v]", result)
	}
}

func TestFindsSubstringIsEqualsOriginalWhenSinglePairParenthesesIsPresent(t *testing.T) {
	s := "()"
	result := longestValidParentheses(s)
	if expected := 2; expected != result {
		t.Errorf("Result must be equals for length [%v] of string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsSubstringLengthWithoutExtraBracket(t *testing.T) {
	s := "())"
	result := longestValidParentheses(s)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsNoSubstringForNotWellFormedString(t *testing.T) {
	s := ")("
	result := longestValidParentheses(s)
	if 0 != result {
		t.Errorf("Result must be zero for not well-formed string [%v], actual is [%v]", s, result)
	}
}

func TestFindsSubstringForRepeatedWellFormedParenthesesWhenIsSequenced(t *testing.T) {
	s := "()()()"
	result := longestValidParentheses(s)
	if expected := 6; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsSubstringForRepeatedWellFormedParenthesesWhenIsNested(t *testing.T) {
	s := "(())"
	result := longestValidParentheses(s)
	if expected := 4; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsLastSubstringWhenStringContainsTwoWellFormedSubstringOnly(t *testing.T) {
	s := "())(())"
	result := longestValidParentheses(s)
	if expected := 4; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsSingleSubstringWhenStartsIsInvalid(t *testing.T) {
	s := "(()"
	result := longestValidParentheses(s)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}
