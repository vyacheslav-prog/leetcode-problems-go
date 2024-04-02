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
	if 0 == len(p) {
		return newPattern(endPattern)
	}
	if 1 < len(p) && zeroOrMoreCharSymbol == p[1] {
		return newZeroOrMorePattern(p[0])
	}
	if anyCharSymbol == p[0] {
		return newPattern(anyCharPattern)
	}
	return newCharPattern(p[0])
}

func parseStringPattern(p string) []pattern {
	nextPattern := detectNextPattern(p)
	if endPattern == nextPattern.name {
		return []pattern{}
	}
	result := []pattern{nextPattern}
	if zeroOrMoreCharPattern == nextPattern.name {
		result = append(result, parseStringPattern(p[2:])...)
	} else if anyCharPattern == nextPattern.name || charPattern == nextPattern.name {
		result = append(result, parseStringPattern(p[1:])...)
	}
	return result
}

func isMatch(s string, p string) bool {
	return p == s
}
