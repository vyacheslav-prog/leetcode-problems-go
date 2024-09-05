package findtheindexofthefirstoccurrenceinastring

func strStr(haystack string, needle string) int {
	var matches []int
	for index := 0; len(haystack) != index; index += 1 {
		matches = append(matches, len(needle)-1)
	}
	if nil == matches {
		return len(needle) - 1
	}
	return matches[len(haystack)-1]
}
