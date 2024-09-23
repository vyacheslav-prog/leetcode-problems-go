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
		for index := sIndex; len(words) != len(wordsOccurences) && sIndex+substringLength != index; index += wordLength {
			for wordIndex := 0; len(words) != wordIndex; wordIndex += 1 {
				if _, isOccured := wordsOccurences[wordIndex]; true != isOccured && s[index:index+wordLength] == words[wordIndex] {
					wordsOccurences[wordIndex] = true
				}
			}
		}
		if 0 != len(wordsOccurences) && len(words) == len(wordsOccurences) {
			indicies = append(indicies, sIndex)
		}
	}
	return indicies
}
