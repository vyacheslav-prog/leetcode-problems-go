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
	var repeatedChars string
	for index, char := range s {
		if index != 0 && byte(char) != repeatedChars[index-1] {
			break
		}
		repeatedChars = s[:index+1]
	}
	return repeatedChars
}
