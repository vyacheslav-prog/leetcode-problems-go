package regularexpressionmatching

const (
	anyCharSymbol        = '.'
	zeroOrMoreCharSymbol = '*'
)

const (
	anyCharPattern        = "anychar"
	charPattern           = "char"
	endPattern            = "endpattern"
	zeroOrMoreCharPattern = "zeroormorechar"
)

type pattern struct {
	length                  int
	name                    string
	zeroOrMorePrecedingChar byte
}

func (p *pattern) is(candidate pattern) bool {
	return p.zeroOrMorePrecedingChar == candidate.zeroOrMorePrecedingChar
}

func (p *pattern) match(s string) (int, bool) {
	switch p.name {
	case anyCharPattern:
		return matchAnyCharPattern(s)
	case charPattern:
		return matchCharPattern(p.zeroOrMorePrecedingChar, s)
	case zeroOrMoreCharPattern:
		return matchZeroOrMoreCharPattern(p.zeroOrMorePrecedingChar, s)
	}
	return 0, false
}

func newAnyCharPattern() pattern {
	return pattern{1, anyCharPattern, 0}
}

func newCharPattern(char byte) pattern {
	return pattern{1, charPattern, char}
}

func newZeroOrMorePattern(precedingChar byte) pattern {
	return pattern{2, zeroOrMoreCharPattern, precedingChar}
}

func detectNextPattern(p string) pattern {
	if 0 == len(p) {
		return pattern{0, endPattern, 0}
	}
	if 1 < len(p) && zeroOrMoreCharSymbol == p[1] {
		return newZeroOrMorePattern(p[0])
	}
	if anyCharSymbol == p[0] {
		return newAnyCharPattern()
	}
	return newCharPattern(p[0])
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

func matchZeroOrMoreCharPattern(c byte, s string) (int, bool) {
	var length int
	for index := 0; index != len(s) && c == s[index]; index += 1 {
		length += 1
	}
	return length, true
}

func parseStringPattern(p string) []pattern {
	nextPattern := detectNextPattern(p)
	if 0 == nextPattern.length {
		return []pattern{}
	}
	return append([]pattern{nextPattern}, parseStringPattern(p[nextPattern.length:])...)
}

func isMatch(s string, p string) bool {
	patterns := parseStringPattern(p)
	if 1 == len(patterns) {
		_, result := patterns[0].match(s)
		return result
	}
	return p == s
}
