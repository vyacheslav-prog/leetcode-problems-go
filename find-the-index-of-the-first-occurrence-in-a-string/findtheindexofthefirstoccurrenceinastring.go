package findtheindexofthefirstoccurrenceinastring

func strStr(haystack string, needle string) int {
	var matches []int
	for index := 0; len(haystack) != index; index += 1 {
		matches = append(matches, len(needle)-1)
	}
	for _, needleChar := range needle {
		for haystackIndex, haystackChar := range haystack {
			if needleChar == haystackChar {
				matches = append(matches, haystackIndex)
			} else {
				matches = append(matches, -1)
			}
		}
	}
	if nil == matches {
		return -1
	}
	return matches[len(matches)-1]
}
