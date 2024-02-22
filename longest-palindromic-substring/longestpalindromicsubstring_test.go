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

func TestFindsSingleLengthSubstringForTwoDifferenceChars(t *testing.T) {
	s := "12"
	result := longestPalindrome(s)
	if result != "1" && result != "2" {
		t.Errorf("Result must be single length substring for %v, result %v", s, result)
	}
}

func TestFindsSelfForSingleRepeatedChar(t *testing.T) {
	s := "xxxx"
	result := longestPalindrome(s)
	if result != s {
		t.Errorf("Result must be original string for %v, actual: %v", s, result)
	}
}

func TestFindsLongestRepeatedSubstringWhenHasOtherChar(t *testing.T) {
	s := "aac"
	result := longestPalindrome(s)
	if result != "aa" {
		t.Errorf("Result must be longest repeated substring for %v, actual %v", s, result)
	}
}

func TestDetectsPalindromeForEmptyString(t *testing.T) {
	result := isPalindrome("")
	if result != true {
		t.Error("Empty string is always palindrome")
	}
}

func TestDetectsNoPalindromeForTwoOtherChars(t *testing.T) {
	s := "qw"
	result := isPalindrome(s)
	if result != false {
		t.Errorf("Two other chars [%v] must detects as palindrome", s)
	}
}
