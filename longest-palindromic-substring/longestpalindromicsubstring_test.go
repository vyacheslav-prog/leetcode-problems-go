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

func TestDetectsNoPalindromeForTwoDifferentChars(t *testing.T) {
	s := "qw"
	result := isPalindrome(s)
	if result != false {
		t.Errorf("Two different chars [%v] must no detects as palindrome", s)
	}
}

func TestDetectsPalindromeForMirroredOddChars(t *testing.T) {
	s := "qwq"
	result := isPalindrome(s)
	if result != true {
		t.Errorf("Mirrored odd chard [%v] must defined as palindrome", s)
	}
}

func TestDetectsPalindromeForRepeatedEvenChars(t *testing.T) {
	s := "8888"
	result := isPalindrome(s)
	if result != true {
		t.Errorf("Repeated char for event counts [%v] must defined as palindrome", s)
	}
}

func TestDetectsNoPalindromeForDifferentChars(t *testing.T) {
	s := "1234"
	result := isPalindrome(s)
	if result != false {
		t.Errorf("Different chars [%v] must detects as no palindrome", s)
	}
}

func TestFindsLongestPalindromeForRepeatedCharWhenFirstIsOther(t *testing.T) {
	s := "mooo"
	result := longestPalindrome(s)
	if result != "ooo" {
		t.Errorf("Result must be [ooo] for string [%v], actual [%v]", s, result)
	}
}
