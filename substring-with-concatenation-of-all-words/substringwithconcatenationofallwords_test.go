package substringwithconcatenationofallwords

import "testing"

func TestFindsNoIndiciesForEmptyString(t *testing.T) {
	result := findSubstring("", nil)
	if 0 != len(result) {
		t.Errorf("Result must be zero for empty string, actual is [%v]", result)
	}
}

func TestFindsZeroIndexForWordEqualsToString(t *testing.T) {
	s, words := "a", []string{"a"}
	result := findSubstring(s, words)
	if expected := [1]int{0}; 1 != len(result) || expected != [1]int(result) {
		t.Errorf("Result must be [%v] for string [%v] and words [%v], actual is [%v]", expected, s, words, result)
	}
}
