package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	hasRepeatedChars := true
	previousChar := s[0]
	for _, value := range s {
		if previousChar != byte(value) {
			hasRepeatedChars = false
		}
	}
	if hasRepeatedChars != true {
		return length
	}
	return 1
}
