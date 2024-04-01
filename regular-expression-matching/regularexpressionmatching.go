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

func newCharPattern(char byte) pattern {
	return pattern{charPattern, char}
}

func newPattern(name string) pattern {
	return pattern{name, 0}
}

func newZeroOrMorePattern(precedingChar byte) pattern {
	return pattern{zeroOrMoreCharPattern, precedingChar}
}

func detectNextPattern(p string) pattern {
	if 1 < len(p) && zeroOrMoreCharSymbol == p[1] {
		return newZeroOrMorePattern(p[0])
	}
	if anyCharSymbol == p[0] {
		return newPattern(anyCharPattern)
	}
	return newCharPattern(p[0])
}

func parseStringPattern(p string) []pattern {
	if 2 == len(p) {
		return []pattern{detectNextPattern(p[:1]), detectNextPattern(p[1:])}
	}
	if "" != p {
		return []pattern{detectNextPattern(p)}
	}
	return []pattern{}
}

func isMatch(s string, p string) bool {
	return p == s
}
