package regularexpressionmatching

const (
	anyCharSymbol        = '.'
	zeroOrMoreCharSymbol = '*'
)

const (
	anyCharPattern        = "anychar"
	charPattern           = "char"
	zeroOrMoreCharPattern = "zeroormorechar"
)

type pattern struct {
	name                    string
	zeroOrMorePrecedingChar byte
}

func (p *pattern) is(candidate pattern) bool {
	return p.zeroOrMorePrecedingChar == candidate.zeroOrMorePrecedingChar
}

func newPattern(name string) pattern {
	return pattern{name, 0}
}

func newZeroOrMorePattern(precedingChar byte) pattern {
	return pattern{zeroOrMoreCharPattern, precedingChar}
}

func detectNextPattern(p string) pattern {
	if 1 < len(p) && zeroOrMoreCharSymbol == p[1] {
		return newPattern(zeroOrMoreCharPattern)
	}
	if anyCharSymbol == p[0] {
		return newPattern(anyCharPattern)
	}
	return newPattern(charPattern)
}

func isMatch(s string, p string) bool {
	return p == s
}
