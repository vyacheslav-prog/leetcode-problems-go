package regularexpressionmatching

const (
	anyCharSymbol        = '.'
	zeroOrMoreCharSymbol = '*'
)

const (
	anyCharPatternName        = "anychar"
	charPatternName           = "char"
	endPatternName            = "endpattern"
	zeroOrMoreCharPatternName = "zeroormorechar"
)

type patternInterface interface {
	match(string) (int, bool)
}

type pattern struct {
	name                    string
	zeroOrMorePrecedingChar byte
}

type anyCharPattern struct{}

func (p anyCharPattern) match(s string) (int, bool) {
	return 1, len(s) != 0
}

type charPattern struct {
	precedingChar byte
}

func (p charPattern) match(s string) (int, bool) {
	if 0 != len(s) && p.precedingChar == s[0] {
		return 1, true
	}
	return 0, false
}

type endPattern struct{}

func (p endPattern) match(s string) (int, bool) {
	return 0, len(s) == 0
}

type zeroOrMoreCharPattern struct {
	zeroOrMorePrecedingChar byte
}

func (p zeroOrMoreCharPattern) match(s string) (int, bool) {
	if anyCharSymbol == p.zeroOrMorePrecedingChar {
		return len(s), true
	}
	var length int
	for length != len(s) && s[length] == p.zeroOrMorePrecedingChar {
		length += 1
	}
	return length, true
}

func (p pattern) match(s string) (int, bool) {
	switch p.name {
	case anyCharPatternName:
		return matchAnyCharPattern(s)
	case charPatternName:
		return matchCharPattern(p.zeroOrMorePrecedingChar, s)
	case endPatternName:
		return matchEndPattern(s)
	case zeroOrMoreCharPatternName:
		return matchZeroOrMoreCharPattern(p.zeroOrMorePrecedingChar, s)
	}
	return 0, false
}

func newAnyCharPattern() pattern {
	return pattern{anyCharPatternName, 0}
}

func newCharPattern(char byte) pattern {
	return pattern{charPatternName, char}
}

func newEndPattern() pattern {
	return pattern{endPatternName, 0}
}

func newZeroOrMorePattern(precedingChar byte) pattern {
	return pattern{zeroOrMoreCharPatternName, precedingChar}
}

func detectNextPattern(p string) (int, patternInterface) {
	if 0 == len(p) {
		return 0, endPattern{}
	}
	if 1 < len(p) && zeroOrMoreCharSymbol == p[1] {
		return 2, zeroOrMoreCharPattern{p[0]}
	}
	if anyCharSymbol == p[0] {
		return 1, anyCharPattern{}
	}
	return 1, charPattern{p[0]}
}

func matchAnyCharPattern(s string) (int, bool) {
	return 1, len(s) != 0
}

func matchCharPattern(c byte, s string) (int, bool) {
	if 0 != len(s) && c == s[0] {
		return 1, true
	}
	return 0, false
}

func matchEndPattern(s string) (int, bool) {
	return 0, len(s) == 0
}

func matchZeroOrMoreCharPattern(c byte, s string) (int, bool) {
	if anyCharSymbol == c {
		return len(s), true
	}
	var length int
	for length != len(s) && s[length] == c {
		length += 1
	}
	return length, true
}

func parseStringPattern(p string) []patternInterface {
	if length, nextPattern := detectNextPattern(p); length != 0 {
		return append([]patternInterface{nextPattern}, parseStringPattern(p[length:])...)
	}
	return []patternInterface{}
}

func isMatch(s string, p string) bool {
	var result bool
	for _, pattern := range append(parseStringPattern(p), endPattern{}) {
		capturedLength, isMatched := pattern.match(s)
		result = isMatched
		if result != true {
			break
		}
		s = s[capturedLength:]
	}
	return result
}
