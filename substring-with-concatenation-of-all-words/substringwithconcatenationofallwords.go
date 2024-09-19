package substringwithconcatenationofallwords

func findSubstring(s string, words []string) []int {
	var indicies []int
	for sIndex := 0; len(s) != sIndex; sIndex += 1 {
		if words[0] == s[sIndex:] {
			indicies = append(indicies, sIndex)
		}
	}
	return indicies
}
