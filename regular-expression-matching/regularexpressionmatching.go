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
	name string
}

func (p *pattern) is(candidate pattern) bool {
	return true
}

func detectNextPattern(p string) pattern {
	if 1 < len(p) && zeroOrMoreCharSymbol == p[1] {
		return pattern{zeroOrMoreCharPattern}
	}
	if anyCharSymbol == p[0] {
		return pattern{anyCharPattern}
	}
	return pattern{charPattern}
}

func isMatch(s string, p string) bool {
	return p == s
}
