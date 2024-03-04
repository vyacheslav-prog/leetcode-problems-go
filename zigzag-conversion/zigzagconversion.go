package zigzagconversion

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	var result string
	var offset int
	for len(result) != len(s) {
		result += string(s[offset])
		offset += numRows
		if len(s) <= offset {
			offset = len(s) - (len(s) - len(result))
		}
	}
	return result
}
