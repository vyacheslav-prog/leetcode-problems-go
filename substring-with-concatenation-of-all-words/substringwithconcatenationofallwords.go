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
		for searchIndex := sIndex; sIndex+substringLength != searchIndex; searchIndex += wordLength {
			sChunks = append(sChunks, s[searchIndex:searchIndex+wordLength])
		}
		for wordIndex, word := range words {
			foundIndex := -1
			for chunkIndex := 0; len(sChunks) != chunkIndex && -1 == foundIndex; chunkIndex += 1 {
				if _, isOccured := wordsOccurences[wordIndex]; true != isOccured && word == sChunks[chunkIndex] {
					foundIndex = wordIndex
				}
			}
			if -1 == foundIndex {
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
