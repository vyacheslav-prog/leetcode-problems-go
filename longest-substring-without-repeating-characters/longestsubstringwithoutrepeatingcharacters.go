package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	currentSubstringLength := 0
	maxSubstringLength := 0
	currentSubstringChars := map[rune]int{}
	for _, value := range s {
		if _, alreadyExists := currentSubstringChars[value]; alreadyExists {
			if maxSubstringLength < currentSubstringLength {
				maxSubstringLength = currentSubstringLength
			}
			currentSubstringChars = map[rune]int{}
			currentSubstringLength = 1
		} else {
			currentSubstringLength += 1
		}
		currentSubstringChars[value] = 1
	}
	if maxSubstringLength < currentSubstringLength {
		return currentSubstringLength
	}
	return maxSubstringLength
}
