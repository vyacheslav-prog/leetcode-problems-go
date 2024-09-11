package findtheindexofthefirstoccurrenceinastring

func strStr(haystack string, needle string) int {
	var matches []int
	for index := 0; len(haystack) != index; index += 1 {
		matches = append(matches, -1)
	}
	for needleIndex, needleChar := range needle {
		needleOffsetOfMatch := len(haystack) * (needleIndex + 1)
		for haystackIndex, haystackChar := range haystack {
			currentMatch := -1
			if prevMatch := matches[needleOffsetOfMatch+haystackIndex-1]; needleIndex == haystackIndex {
				if needleChar != haystackChar {
					currentMatch = -1
				} else if -1 != prevMatch {
					currentMatch = prevMatch
				} else {
					currentMatch = haystackIndex
				}
			} else if needleChar < haystackChar {
				currentMatch = prevMatch
			}
			matches = append(matches, currentMatch)
		}
	}
	if nil == matches {
		return -1
	}
	return matches[len(matches)-1]
}
