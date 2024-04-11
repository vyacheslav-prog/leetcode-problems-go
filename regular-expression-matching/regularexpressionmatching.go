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
	name                    string
	zeroOrMorePrecedingChar byte
}

func (p *pattern) match(s string) (int, bool) {
	switch p.name {
	case anyCharPattern:
		return matchAnyCharPattern(s)
	case charPattern:
		return matchCharPattern(p.zeroOrMorePrecedingChar, s)
	case endPattern:
		return matchEndPattern(s)
	case zeroOrMoreCharPattern:
		return matchZeroOrMoreCharPattern(p.zeroOrMorePrecedingChar, s)
	}
	return 0, false
}

func newAnyCharPattern() pattern {
	return pattern{anyCharPattern, 0}
}

func newCharPattern(char byte) pattern {
	return pattern{charPattern, char}
}

func newEndPattern() pattern {
	return pattern{endPattern, 0}
}

func newZeroOrMorePattern(precedingChar byte) pattern {
	return pattern{zeroOrMoreCharPattern, precedingChar}
}

func detectNextPattern(p string) (int, pattern) {
	if 0 == len(p) {
		return 0, newEndPattern()
	}
	if 1 < len(p) && zeroOrMoreCharSymbol == p[1] {
		return 2, newZeroOrMorePattern(p[0])
	}
	if anyCharSymbol == p[0] {
		return 1, newAnyCharPattern()
	}
	return 1, newCharPattern(p[0])
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

func parseStringPattern(p string) []pattern {
	length, nextPattern := detectNextPattern(p)
	if endPattern == nextPattern.name {
		return []pattern{}
	}
	return append([]pattern{nextPattern}, parseStringPattern(p[length:])...)
}

func isMatch(s string, p string) bool {
	var result bool
	for _, pattern := range append(parseStringPattern(p), newEndPattern()) {
		capturedLength, isMatched := pattern.match(s)
		result = isMatched
		if result != true {
			break
		}
		s = s[capturedLength:]
	}
	return result
}
