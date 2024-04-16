package regularexpressionmatching

const (
	anyCharSymbol        = '.'
	zeroOrMoreCharSymbol = '*'
)

type pattern interface {
	match(string) (int, bool)
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

type zeroOrMoreAnyCharPattern struct {
	terminator pattern
}

func (p zeroOrMoreAnyCharPattern) match(s string) (int, bool) {
	return len(s), true
}

type zeroOrMoreCharPattern struct {
	zeroOrMorePrecedingChar byte
}

func (p zeroOrMoreCharPattern) match(s string) (int, bool) {
	var length int
	for length != len(s) && s[length] == p.zeroOrMorePrecedingChar {
		length += 1
	}
	return length, true
}

func detectNextPattern(p string) (int, pattern) {
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

func parseStringPattern(p string) []pattern {
	if length, nextPattern := detectNextPattern(p); length != 0 {
		return append([]pattern{nextPattern}, parseStringPattern(p[length:])...)
	}
	return []pattern{}
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
