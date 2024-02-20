package longestpalindromicsubstring

import "testing"

func TestFindsSelfWhenStringIsEmpty(t *testing.T) {
	result := longestPalindrome("")
	if result != "" {
		t.Error("Result must be an empty string when input string is empty")
	}
}

func TestFindsSelfWhenStringIsOneChar(t *testing.T) {
	s := "a"
	result := longestPalindrome(s)
	if result != s {
		t.Errorf("Result must be original string for one char string, result: %v", result)
	}
}
