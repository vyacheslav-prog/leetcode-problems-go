package substringwithconcatenationofallwords

import "testing"

func TestFindsNoIndiciesForEmptyString(t *testing.T) {
	result := findSubstring("", nil)
	if 0 != len(result) {
		t.Errorf("Result must be zero for empty string, actual is [%v]", result)
	}
}

func TestFindsZeroIndexForSingleCharWordEqualsToString(t *testing.T) {
	s, words := "a", []string{"a"}
	result := findSubstring(s, words)
	if expected := [1]int{0}; 1 != len(result) || expected != [1]int(result) {
		t.Errorf("Result must be [%v] for string [%v] and words [%v], actual is [%v]", expected, s, words, result)
	}
}

func TestFindsNoIndiciesForWordsWithoutSingleCharString(t *testing.T) {
	s, words := "b", []string{"c"}
	result := findSubstring(s, words)
	if 0 != len(result) {
		t.Errorf("Result must be empty for string [%v] is not contained in words [%v], actual is [%v]", s, words, result)
	}
}

func TestFindsNoIndicesForEmptyWordsAndNotEmptyString(t *testing.T) {
	s, words := "bob", []string{}
	result := findSubstring(s, words)
	if 0 != len(result) {
		t.Errorf("Result must be empty for string [%v] and empty words [%v], actual is [%v]", s, words, result)
	}
}

func TestFindsTwoIndiciesForRepeatedWordIntoString(t *testing.T) {
	s, words := "zz", []string{"z"}
	result := findSubstring(s, words)
	if expected := [2]int{0, 1}; 2 != len(result) || expected != [2]int(result) {
		t.Errorf("Result must be [%v] for string [%v] and words [%v], actual is [%v]", expected, s, words, result)
	}
}
