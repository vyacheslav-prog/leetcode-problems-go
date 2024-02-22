package longestpalindromicsubstring

func isPalindrome(s string) bool {
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
