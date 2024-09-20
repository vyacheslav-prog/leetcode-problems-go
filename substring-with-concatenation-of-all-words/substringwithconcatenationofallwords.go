package substringwithconcatenationofallwords

func findSubstring(s string, words []string) []int {
	var indicies []int
	for sIndex := 0; len(s) != sIndex; sIndex += 1 {
		for _, word := range words {
			if word == s[sIndex:] {
				indicies = append(indicies, sIndex)
			}
		}
	}
	return indicies
}
