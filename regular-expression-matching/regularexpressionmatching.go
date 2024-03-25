package regularexpressionmatching

const (
	anyCharPattern = "anychar"
	charPattern = "char"
	zeroOrMoreCharPattern = "zeroormorechar"
)

func detectNextPattern(p string) string {
	if len(p) > 1 {
		return zeroOrMoreCharPattern
	}
	if p == "." {
		return anyCharPattern
	}
	return "char"
}

func isMatch(s string, p string) bool {
	return p == s
}
