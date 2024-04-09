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
	if result.is(newAnyCharPattern()) != true {
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
	if result.is(newAnyCharPattern()) != true {
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
	if len(result) != 1 || result[0].is(newAnyCharPattern()) != true {
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

func TestMismatchesCharPatternForEmptyString(t *testing.T) {
	var c byte = 'a'
	_, result := matchCharPattern(c, "")
	if result != false {
		t.Errorf("Empty string must not matches [false] a char [%v], actual is [%v]", c, result)
	}
}

func TestMatchesCharPatternWithEqualsSingleCharString(t *testing.T) {
	var c byte = 'a'
	s := "a"
	length, result := matchCharPattern(c, s)
	if result != true || length != 1 {
		t.Errorf("Result must be [true] for match [%v] into string [%v], actual is [%v] and [%v]", c, s, length, result)
	}
}

func TestMismatchesCharPatternForAnotherCharIntoString(t *testing.T) {
	var c byte = 'b'
	s := "c"
	_, result := matchCharPattern(c, s)
	if result != false {
		t.Errorf("Result must be [false] for char [%v] for string [%v], actual is [%v]", c, s, result)
	}
}

func TestMismatchesAnyCharPatternForEmptyString(t *testing.T) {
	_, result := matchAnyCharPattern("")
	if result != false {
		t.Errorf("Result must be [false] for match any char pattern for empty string, actual is [%v]", result)
	}
}

func TestMatchesAnyCharPatternForNotEmptyString(t *testing.T) {
	s := "a"
	length, result := matchAnyCharPattern(s)
	if result != true || length != 1 {
		t.Errorf("Result must be [1] and [true] for any char pattern for string [%v], actual is [%v] and [%v]", s, length, result)
	}
}

func TestMatchesZeroOrMoreCharPatternForEmptyString(t *testing.T) {
	var c byte = 'a'
	length, result := matchZeroOrMoreCharPattern(c, "")
	if result != true || length != 0 {
		t.Errorf("Result must be [0] and [true] for zero or more char [%v] pattern empty string, actual is [%v] and [%v]", c, length, result)
	}
}

func TestMatchesZeroOrMoreCharPatternForSuitableSingleCharString(t *testing.T) {
	var c byte = 'a'
	s := "a"
	length, result := matchZeroOrMoreCharPattern(c, s)
	if result != true || length != 1 {
		t.Errorf("Result must be [1] and [true] for zero or more [%v] for string [%s], actual is [%v] and [%v]", c, s, length, result)
	}
}

func TestMatchesZeroOrMoreCharPatternForNotSuitableSingleCharString(t *testing.T) {
	var c byte = 'a'
	s := "b"
	length, result := matchZeroOrMoreCharPattern(c, s)
	if result != true || length != 0 {
		t.Errorf("Result must be [0] and [true] for zero or more [%v] for string [%s], actual is [%v] and [%v]", c, s, length, result)
	}
}

func TestMatchesZeroOrMoreCharPatternForRepeatedSuitableChars(t *testing.T) {
	var c byte = 'a'
	s := "aa"
	length, result := matchZeroOrMoreCharPattern(c, s)
	if result != true || length != 2 {
		t.Errorf("Result must be [2] and [true] for zero or more [%v] for string [%s], actual is [%v] and [%v]", c, s, length, result)
	}
}

func TestMatchesZeroOrMoreCharPatternForNotOnlyRepeatedAndSuitableChars(t *testing.T) {
	var c byte = 'a'
	s := "aab"
	length, result := matchZeroOrMoreCharPattern(c, s)
	if result != true || length != 2 {
		t.Errorf("Result must be [2] and [true] for zero or more [%v] for string [%s], actual is [%v] and [%v]", c, s, length, result)
	}
}

func TestPassesAnyCharPatternForSingleCharString(t *testing.T) {
	s, p := "a", "."
	result := isMatch(s, p)
	if result != true {
		t.Errorf("Result must be [true] for string [%v] and any char pattern [%v], actual is [%v]", s, p, result)
	}
}

func TestPassesZeroOrMoreCharPatternForEmptyString(t *testing.T) {
	p := "a*"
	result := isMatch("", p)
	if result != true {
		t.Errorf("Result must be [true] for empty string for zero or more char pattern [%v], actual is [%v]", p, result)
	}
}

func TestPassesTwoAnyCharPatternForTwoAnyCharString(t *testing.T) {
	s, p := "ab", ".."
	result := isMatch(s, p)
	if result != true {
		t.Errorf("Result must be [true] for two char [%v] and twice any char pattern [%v], actual is [%v]", s, p, result)
	}
}

func TestRejectsAnyCharAndCharPatternsForTwoDifferentChars(t *testing.T) {
	s, p := "ab", ".c"
	result := isMatch(s, p)
	if result != false {
		t.Errorf("Result must be [false] for two char [%v] and an any char and char pattern [%v], actual is [%v]", s, p, result)
	}
}
