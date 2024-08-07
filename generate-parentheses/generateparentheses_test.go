package generateparentheses

import "testing"

func TestGeneratesNilForZeroPairs(t *testing.T) {
	result := generateParenthesis(0)
	if nil != result {
		t.Errorf("Result must be nil for zero pairs, actual is [%v]", result)
	}
}

func TestGeneratesOnePairParentheses(t *testing.T) {
	n := 1
	result := generateParenthesis(n)
	if expected := "()"; 1 != len(result) || expected != result[0] {
		t.Errorf("Result must have [%v] for pairs n [%v], actual is [%v]", expected, n, result)
	}
}

func TestGeneratesTwoPairParentheses(t *testing.T) {
	n := 2
	result := generateParenthesis(n)
	if expected := [2]string{"(())", "()()"}; 2 != len(result) || expected != [2]string(result) {
		t.Errorf("Result must be [%v] for pairs n [%v], actual is [%v]", expected, n, result)
	}
}

func TestGeneratesThreePairParentheses(t *testing.T) {
	n := 3
	result := generateParenthesis(n)
	if expected := [5]string{"((()))", "(()())", "(())()", "()(())", "()()()"}; 5 != len(result) || expected != [5]string(result) {
		t.Errorf("Result must be [%v] for pairs n [%v], actual is [%v]", expected, n, result)
	}
}
