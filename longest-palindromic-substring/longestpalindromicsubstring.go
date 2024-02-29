package longestpalindromicsubstring

func isPalindrome(s string) bool {
	leftPointer, rightPointer := 0, len(s)-1
	if rightPointer == -1 {
		return true
	}
	for (rightPointer - leftPointer) > 0 {
		if s[leftPointer] != s[rightPointer] {
			return false
		}
		leftPointer += 1
		rightPointer -= 1
	}
	return true
}

func longestPalindrome(s string) string {
	if isPalindrome(s) {
		return s
	}
	var result string
	for window := len(s); result == ""; window -= 1 {
		for leftOffset := 0; len(s) != (leftOffset + window - 1); leftOffset += 1 {
			candidate := s[leftOffset : leftOffset+window]
			if isPalindrome(candidate) {
				result = candidate
				break
			}
		}
	}
	return result
}
