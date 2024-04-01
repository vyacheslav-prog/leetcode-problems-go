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
	if expected := newCharPattern('a'); result.is(expected) != true {
		t.Errorf("Result must be [%v] for detected pattern [%v], actual is [%v]", expected, p, result)
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
	if expected := newZeroOrMorePattern('a'); result.is(expected) != true {
		t.Errorf("Result must be [%v] for detected pattern [%v], actual is [%v]", expected, p, result)
	}
}

func TestDetectsCharPatternForTwoCharString(t *testing.T) {
	p := "ab"
	result := detectNextPattern(p)
	if expected := newCharPattern('a'); result.is(expected) != true {
		t.Errorf("Result must be [%v] for detected pattern [%v] with two char, actual is [%v]", expected, p, result)
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
	if expected := newZeroOrMorePattern('a'); result.is(expected) != true {
		t.Errorf("Result must be [%v] for pattern [%v] precending a char, actual is [%v]", expected, p, result)
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
	if expected := newCharPattern('a'); len(result) != 1 || result[0].is(expected) != true {
		t.Errorf("Result must have [%v] pattern for string [%v], actual is [%v]", expected, p, result)
	}
}

func TestParsesPatternWithSingleAnyChar(t *testing.T) {
	p := "."
	result := parseStringPattern(p)
	if len(result) != 1 || result[0].is(newPattern(anyCharPattern)) != true {
		t.Errorf("Result must have [anychar] pattern for string [%v], actual is [%v]", p, result)
	}
}

func TestParsesTwoCharPatternsFromString(t *testing.T) {
	p := "aa"
	result := parseStringPattern(p)
	if expected := newCharPattern('a'); len(result) != 2 || result[0].is(expected) != true || result[1].is(expected) != true {
		t.Errorf("Result must have two [%v] pattern for [%v], actual is [%v]", expected, p, result)
	}
}

func TestParsesCharAndZeroOrMoreCharPatternFromString(t *testing.T) {
	p := "a*b"
	result := parseStringPattern(p)
	if expected := []pattern{newZeroOrMorePattern('a'), newCharPattern('b')}; 2 != len(result) || result[0].is(expected[0]) != true || result[1].is(expected[1]) != true {
		t.Errorf("Result must be [%v] for pattern string [%v], actual is [%v]", expected, p, result)
	}
}
