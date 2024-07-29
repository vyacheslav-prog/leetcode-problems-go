package generateparentheses

import "testing"

func TestGeneratesNilForZeroPairs(t *testing.T) {
	result := generateParentheses(0)
	if nil != result {
		t.Errorf("Result must be nil for zero pairs, actual is [%v]", result)
	}
}

func TestGeneratesOnePairParentheses(t *testing.T) {
	n := 1
	result := generateParentheses(n)
	if expected := "()"; 1 != len(result) || expected != result[0] {
		t.Errorf("Result must have [%v] for pairs n [%v], actual is [%v]", expected, n, result)
	}
}
