package substringwithconcatenationofallwords

func findSubstring(s string, words []string) []int {
	var indicies []int
	var sChunks []string
	var substringLength, wordLength int
	var wordsOccurences map[int]bool
	if 0 != len(words) {
		wordLength = len(words[0])
		substringLength = len(words) * wordLength
	}
	for sIndex := 0; len(s)-substringLength+1 != sIndex; sIndex += 1 {
		sChunks = nil
		wordsOccurences = make(map[int]bool)
		for chunkIndex := sIndex; sIndex+substringLength != chunkIndex; chunkIndex += wordLength {
			sChunks = append(sChunks, s[chunkIndex:chunkIndex+wordLength])
		}
		for _, word := range words {
			foundIndex := 0
			for len(sChunks) != foundIndex {
				if _, isOccured := wordsOccurences[foundIndex]; true != isOccured && word == sChunks[foundIndex] {
					break
				}
				foundIndex += 1
			}
			if len(sChunks) == foundIndex {
				break
			}
			wordsOccurences[foundIndex] = true
		}
		if 0 != len(wordsOccurences) && len(words) == len(wordsOccurences) {
			indicies = append(indicies, sIndex)
		}
	}
	return indicies
}
