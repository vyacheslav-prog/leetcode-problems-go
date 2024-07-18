package validparentheses

import "testing"

func TestPassesEmptyString(t *testing.T) {
	result := isValid("")
	if true != result {
		t.Errorf("Result must be valid for empty string, actual is [%v]", result)
	}
}

func TestRejectsOneBracketOnly(t *testing.T) {
	s := "("
	result := isValid(s)
	if false != result {
		t.Errorf("Result must be [false] for single bracket string [%v], actual is [%v]", s, result)
	}
}

func TestPassesTwoBalancedBrackets(t *testing.T) {
	s := "[]"
	result := isValid(s)
	if true != result {
		t.Errorf("Result must be [true] for balanced brackets into string [%v], actual is [%v]", s, result)
	}
}

func TestRejectsUnbalancedBrackets(t *testing.T) {
	s := "((]"
	result := isValid(s)
	if false != result {
		t.Errorf("Result must be [false] for unbalanced brackets into string [%v], actual is [%v]", s, result)
	}
}

func TestPassesBalancedAllBracketPairsIntoSequence(t *testing.T) {
	s := "()[]{}"
	result := isValid(s)
	if true != result {
		t.Errorf("Result must be [true] for sequence with all balanced brackets into string [%v], actual is [%v]", s, result)
	}
}

func TestRejectsCloseToOpenSameBrackets(t *testing.T) {
	s := ")("
	result := isValid(s)
	if false != result {
		t.Errorf("Result must be [false] for close-to-open same brackets into string [%v], actual is [%v]", s, result)
	}
}
