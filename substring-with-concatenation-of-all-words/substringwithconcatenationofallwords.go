package substringwithconcatenationofallwords

func findSubstring(s string, words []string) []int {
	var indicies []int
	var substringLength, wordLength int
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordLength = len(word)
		substringLength += wordLength
		wordsMap[word] += 1
	}
	for sIndex := 0; len(s)-substringLength+1 != sIndex; sIndex += 1 {
		foundWords, chunkIndex := make(map[string]int), sIndex
		for ; sIndex+substringLength != chunkIndex; chunkIndex += wordLength {
			chunk := s[chunkIndex : chunkIndex+wordLength]
			if _, isOccured := wordsMap[chunk]; true != isOccured {
				break
			}
			foundWords[chunk] += 1
			if wordsMap[chunk] < foundWords[chunk] {
				break
			}
		}
		if 0 != len(foundWords) && substringLength == chunkIndex-sIndex {
			indicies = append(indicies, sIndex)
		}
	}
	return indicies
}
