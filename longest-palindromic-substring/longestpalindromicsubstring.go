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
	toLeft, toRight := longestPalindrome(s[0:len(s)-1]), longestPalindrome(s[1:len(s)])
	if len(toLeft) > len(toRight) {
		return toLeft
	}
	return toRight
}
