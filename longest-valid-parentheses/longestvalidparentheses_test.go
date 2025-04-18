package longestvalidparentheses

import "testing"

func BenchmarkFindsSubstringForAnyThousandBrackets(b *testing.B) {
	var chars [1000]rune
	for i := 0; i != len(chars); i += 1 {
		if i%6 < 3 || i%10/3 < 1 {
			chars[i] = openingBracket
		} else {
			chars[i] = closingBracket
		}
	}
	s := string(chars[:])
	for i := 0; i < b.N; i += 1 {
		longestValidParentheses(s)
	}
}

func BenchmarkForTwentyClosingBrackets(b *testing.B) {
	var chars [20]rune
	for i := 0; i != len(chars); i += 1 {
		chars[i] = closingBracket
	}
	s := string(chars[:])
	for i := 0; i < b.N; i += 1 {
		longestValidParentheses(s)
	}
}

func BenchmarkForTwentyOpeningBrackets(b *testing.B) {
	var chars [20]rune
	for i := 0; i != len(chars); i += 1 {
		chars[i] = openingBracket
	}
	s := string(chars[:])
	for i := 0; i < b.N; i += 1 {
		longestValidParentheses(s)
	}
}

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

func TestFindsFirstAndLongestSubstringForTwoSubstrings(t *testing.T) {
	s := "()())()"
	result := longestValidParentheses(s)
	if expected := 4; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsSubstringWhenSecondPairsIsNotClosed(t *testing.T) {
	s := "()(()"
	result := longestValidParentheses(s)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for string [%v] with unclosed pair, actual is [%v]", expected, s, result)
	}
}

func TestFindsSubstringWhenUnclosedBracketsAroundParentheses(t *testing.T) {
	s := "(()(()"
	result := longestValidParentheses(s)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsFirstSubstringWhenIsNestedBrackets(t *testing.T) {
	s := "((()))(()"
	result := longestValidParentheses(s)
	if expected := 6; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsSubstringForWellFormedPairsAndExtraOpeningBracketIntoStart(t *testing.T) {
	s := "(()()"
	result := longestValidParentheses(s)
	if expected := 4; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsNestedParenthesesSubstringWhenExtraBracketIsStarted(t *testing.T) {
	s := "((())"
	result := longestValidParentheses(s)
	if expected := 4; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsFirstNestedSubstringWithExtraOpeningBracketIsLast(t *testing.T) {
	s := "(())("
	result := longestValidParentheses(s)
	if expected := 4; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsSingletonLongestSubstringWithNotFirstExtraOpeningBracket(t *testing.T) {
	s := ")(((((()())()()))()(()))("
	result := longestValidParentheses(s)
	if expected := 22; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}

func TestFindsSubstringForExtraOpeningBracketsAround(t *testing.T) {
	s := "(()("
	result := longestValidParentheses(s)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for string [%v], actual is [%v]", expected, s, result)
	}
}
