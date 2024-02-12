package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	currentSubstringLength := 0
	maxSubstringLength := 0
	currentSubstringChars := map[rune]int{}
	for valueIndex, value := range s {
		if currentIndex, alreadyExists := currentSubstringChars[value]; alreadyExists {
			subSubstringLength := lengthOfLongestSubstring(s[currentIndex+1:])
			if subSubstringLength > currentSubstringLength {
				return subSubstringLength
			}
			return currentSubstringLength
		} else {
			currentSubstringLength += 1
		}
		currentSubstringChars[value] = valueIndex
	}
	if maxSubstringLength < currentSubstringLength {
		return currentSubstringLength
	}
	return maxSubstringLength
}
