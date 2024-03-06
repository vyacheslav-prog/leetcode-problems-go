package zigzagconversion

func convert(s string, numRows int) string {
	var result string
	var offset, currentRow int
	periodicOffset := numRows + max(numRows-2, 0)
	hasDiagonalChars := periodicOffset != numRows
	for len(result) != len(s) {
		result += string(s[offset])
		if hasDiagonalChars {
			isNotFirstOrLastRow := currentRow != (numRows-1) && currentRow != 0
			diagonalOffset := offset + periodicOffset - (currentRow * 2)
			if isNotFirstOrLastRow && diagonalOffset < len(s) {
				result += string(s[diagonalOffset])
			}
		}
		offset += periodicOffset
		if len(s) <= offset {
			currentRow += 1
			offset = currentRow
		}
	}
	return result
}
