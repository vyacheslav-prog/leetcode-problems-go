package substringwithconcatenationofallwords

func findSubstring(s string, words []string) []int {
	var indicies []int
	var substringLength, wordLength int
	var wordsOccurences map[int]bool
	if 0 != len(words) {
		wordLength = len(words[0])
		substringLength = len(words) * wordLength
	}
	for sIndex := 0; len(s)-substringLength+1 != sIndex; sIndex += 1 {
		wordsOccurences = make(map[int]bool)
		for searchIndex := sIndex; sIndex+substringLength != searchIndex; searchIndex += wordLength {
			searchWord, wordIndex := s[searchIndex:searchIndex+wordLength], 0
			for len(words) != wordIndex {
				if _, isOccured := wordsOccurences[wordIndex]; true != isOccured && searchWord == words[wordIndex] {
					break
				}
				wordIndex += 1
			}
			if len(words) != wordIndex {
				wordsOccurences[wordIndex] = true
			}
		}
		if 0 != len(wordsOccurences) && len(words) == len(wordsOccurences) {
			indicies = append(indicies, sIndex)
		}
	}
	return indicies
}
