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
	var length int
	if _, tResult := p.terminator.match(s); tResult {
		return 0, false
	} else {
		length += 1
	}
	if len(s) != 0 {
		if _, nResult := p.match(s[1:]); nResult {
			length += 1
		}
	}
	return length, true
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
	firstChar := p[0]
	if 1 == len(p) || zeroOrMoreCharSymbol != p[1] {
		if anyCharSymbol == firstChar {
			return 1, anyCharPattern{}
		}
		return 1, charPattern{firstChar}
	}
	if anyCharSymbol != firstChar {
		return 2, zeroOrMoreCharPattern{firstChar}
	}
	_, nPattern := detectNextPattern(p[2:])
	return 2, zeroOrMoreAnyCharPattern{nPattern}
}

func parseStringPattern(p string) []pattern {
	if length, nextPattern := detectNextPattern(p); length != 0 {
		return append([]pattern{nextPattern}, parseStringPattern(p[length:])...)
	}
	return []pattern{}
}

func planPatterns(numChars int, pl []pattern) []pattern {
	return pl
}

func isMatchBruteForce(s string, p string) bool {
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

func isMatchWithMemo(s string, p string) bool {
	stringRowsNum, patternColumnsNum := len(s)+1, len(p)+1
	solutionsNum := stringRowsNum * patternColumnsNum
	if 1 == solutionsNum {
		return true
	}
	solutions := make([]bool, solutionsNum)
	solutions[0] = true // for empty string and empty pattern
	for patternColumnsCounter := 2; patternColumnsCounter != patternColumnsNum; patternColumnsCounter += 1 {
		if zeroOrMoreCharSymbol == p[patternColumnsCounter-1] {
			solutions[patternColumnsCounter] = solutions[patternColumnsCounter-2]
		}
	}
	for previousStringOffset, stringRowsCounter := 0, 1; stringRowsCounter != stringRowsNum; stringRowsCounter += 1 {
		stringOffset, stringSymbol := previousStringOffset+patternColumnsNum, s[stringRowsCounter-1]
		for patternColumnsCounter := 1; patternColumnsCounter != patternColumnsNum; patternColumnsCounter += 1 {
			patternSymbol, solution := p[patternColumnsCounter-1], solutions[previousStringOffset+patternColumnsCounter-1]
			if zeroOrMoreCharSymbol != patternSymbol {
				solution = solution && (patternSymbol == stringSymbol || anyCharSymbol == patternSymbol)
			} else if prevPatternSymbolSolution := solutions[stringOffset+patternColumnsCounter-2]; prevPatternSymbolSolution {
				solution = prevPatternSymbolSolution
			} else {
				previousPatternSymbol, prevStringSymbolSolution := p[patternColumnsCounter-2], solutions[previousStringOffset+patternColumnsCounter]
				solution = (previousPatternSymbol == stringSymbol || anyCharSymbol == previousPatternSymbol) && prevStringSymbolSolution
			}
			solutions[patternColumnsCounter+stringOffset] = solution
		}
		previousStringOffset += patternColumnsNum
	}
	return solutions[solutionsNum-1]
}

func isMatch(s string, p string) bool {
	return isMatchWithMemo(s, p)
}
