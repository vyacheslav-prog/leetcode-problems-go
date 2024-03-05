package zigzagconversion

func convert(s string, numRows int) string {
	var result string
	var offset, startsRow int
	periodicOffset := numRows + max(numRows-2, 0)
	for len(result) != len(s) {
		result += string(s[offset])
		if periodicOffset != numRows && startsRow != (numRows-1) {
			offset += periodicOffset - startsRow - startsRow
		} else {
			offset += periodicOffset
		}
		if len(s) <= offset {
			startsRow += 1
			offset = startsRow
		}
	}
	return result
}
