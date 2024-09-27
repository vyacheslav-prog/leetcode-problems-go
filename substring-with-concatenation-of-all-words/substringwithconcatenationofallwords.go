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
		for _, word := range words {
			chunkIndex := 0
			for len(sChunks) != chunkIndex {
				if _, isOccured := wordsOccurences[chunkIndex]; true != isOccured && word == sChunks[chunkIndex] {
					break
				}
				chunkIndex += 1
			}
			if len(sChunks) == chunkIndex {
				break
			}
			wordsOccurences[chunkIndex] = true
		}
		if 0 != len(wordsOccurences) && len(words) == len(wordsOccurences) {
			indicies = append(indicies, sIndex)
		}
	}
	return indicies
}
