package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestFindsZeroLengthForEmptyString(t *testing.T) {
	result := lengthOfLongestSubstring("")
	if result != 0 {
		t.Errorf("Result must be a zero for empty string, result: %v", result)
	}
}

func TestFindsSingletonSubstringForOneItemString(t *testing.T) {
	s := "1"
	result := lengthOfLongestSubstring(s)
	if result != 1 {
		t.Errorf("Result must be equals to one for one item string, result: %v", result)
	}
}

func TestFindsOneItemSubstringForTwoSameChars(t *testing.T) {
	s := "aa"
	result := lengthOfLongestSubstring(s)
	if result != 1 {
		t.Errorf("Result must be equals to one for two a same chars the string, result: %v", result)
	}
}

func TestFindsSubstringForTwoOtherChars(t *testing.T) {
	s := "1z"
	result := lengthOfLongestSubstring(s)
	if result != 2 {
		t.Errorf("Result must be equals to two for two a other chars the string, result: %v", result)
	}
}

func TestFindsOneItemSubstringForManySameChars(t *testing.T) {
	s := "          "
	result := lengthOfLongestSubstring(s)
	if result != 1 {
		t.Errorf("Result must be equals to one for many same chars the string, result: %v", result)
	}
}

func TestFindsSubstringForRepeatedChars(t *testing.T) {
	s := "1ab1ab1ab"
	result := lengthOfLongestSubstring(s)
	if result != 3 {
		t.Errorf("Result must be equals a length for repeated chars, result: %v", result)
	}
}

func TestFindsSubstringWhenLongestSubstringIsLast(t *testing.T) {
	s := "aab"
	result := lengthOfLongestSubstring(s)
	if result != 2 {
		t.Errorf("Result must be 2 for \"%v\", when longest substring is last, result: %v", s, result)
	}
}

func TestFindsSubstringWhenFirstCharIsRepeated(t *testing.T) {
	s := "dvdf"
	result := lengthOfLongestSubstring(s)
	if result != 3 {
		t.Errorf("Result must be 3 for \"%v\", when first char is repeated, result: %v", s, result)
	}
}
