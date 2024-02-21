package longestpalindromicsubstring

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var result string
	for index, char := range s {
		result = s[0 : index+1]
		if byte(char) != result[index] {
			return result
		}
	}
	return result
}
