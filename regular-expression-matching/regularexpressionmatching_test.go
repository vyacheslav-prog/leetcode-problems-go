package regularexpressionmatching

import "testing"

func TestPassesForEmptyStringsMatching(t *testing.T) {
	result := isMatch("", "")
	if result != true {
		t.Errorf("Empty string must be [true] matching in empty origin string, actual is [%v]", result)
	}
}

func TestRejectsCharPatternForEmptyString(t *testing.T) {
	s := ""
	p := "a"
	result := isMatch(s, p)
	if result != false {
		t.Errorf("Result must be [false] when a char pattern [%v] for empty string, actual is [%v]", p, result)
	}
}

func TestPassesForEqualsCharPatternAndString(t *testing.T) {
	s := "a"
	p := "a"
	result := isMatch(s, p)
	if result != true {
		t.Errorf("Result must be [true] for equal char pattern [%v] and string [%v], actual is [%v]", p, s, result)
	}
}

func TestDetectsCharPattern(t *testing.T) {
	p := "a"
	result := detectNextPattern(p)
	if result.is(newPattern(charPattern)) != true {
		t.Errorf("Result must be [char] for detected pattern [%v], actual is [%v]", p, result)
	}
}

func TestDetectsAnyCharPattern(t *testing.T) {
	p := "."
	result := detectNextPattern(p)
	if result.is(newPattern(anyCharPattern)) != true {
		t.Errorf("Result must be [anychar] for detected pattern [%v], actual is [%v]", p, result)
	}
}

func TestDetectsZeroOrMoreCharPattern(t *testing.T) {
	p := "a*"
	result := detectNextPattern(p)
	if result.is(newPattern(zeroOrMoreCharPattern)) != true {
		t.Errorf("Result must be [zeroormorechar] for detected pattern [%v], actual is [%v]", p, result)
	}
}

func TestDetectsCharPatternForTwoCharString(t *testing.T) {
	p := "ab"
	result := detectNextPattern(p)
	if result.is(newPattern(charPattern)) != true {
		t.Errorf("Result must be [char] for detected pattern [%v] with two char, actual is [%v]", p, result)
	}
}

func TestDetectsAnyCharPatternPrecendingChar(t *testing.T) {
	p := ".a"
	result := detectNextPattern(p)
	if result.is(newPattern(anyCharPattern)) != true {
		t.Errorf("Result must be [anychar] for an any char pattern [%v] precending a char, actual is [%v]", p, result)
	}
}

func TestDetectsZeroOrMoreCharPatternPrecendingChar(t *testing.T) {
	p := "a*b"
	result := detectNextPattern(p)
	if result.is(newPattern(zeroOrMoreCharPattern)) != true {
		t.Errorf("Result must be [zeroormorechar] for pattern [%v] precending a char, actual is [%v]", p, result)
	}
}

func TestComparesPrecendingCharForZeroOrMoreCharPattern(t *testing.T) {
	p, candidate := newZeroOrMorePattern('a'), newZeroOrMorePattern('b')
	result := p.is(candidate)
	if result != false {
		t.Errorf("Zero or more char pattern [%v] is not equals for [%v], actual is [true]", p, candidate)
	}
}

func TestParsesEmptyStringPattern(t *testing.T) {
	result := parseStringPattern("")
	if len(result) != 0 {
		t.Errorf("Result must be empty list for an empty string pattern, actual length is [%v]", len(result))
	}
}

func TestParsesSinglePatternString(t *testing.T) {
	p := "a"
	result := parseStringPattern(p)
	if len(result) != 1 || result[0].is(newPattern(charPattern)) != true {
		t.Errorf("Result must have [char] pattern for string [%v], actual is [%v]", p, result)
	}
}

func TestParsesPatternWithSingleAnyChar(t *testing.T) {
	p := "."
	result := parseStringPattern(p)
	if len(result) != 1 || result[0].is(newPattern(anyCharPattern)) != true {
		t.Errorf("Result must have [anychar] pattern for string [%v], actual is [%v]", p, result)
	}
}

func TestParsesTwoPatternsFromString(t *testing.T) {
	p := "aa"
	result := parseStringPattern(p)
	if len(result) != 2 || result[0].is(newPattern(charPattern)) != true || result[1].is(newPattern(charPattern)) != true {
		t.Errorf("Result must have two [char] pattern for [%v], actual is [%v]", p, result)
	}
}
