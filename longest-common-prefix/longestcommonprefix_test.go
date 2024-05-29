package longestcommonprefix

import "testing"

func TestFindsEmptyPrefixForZeroStrings(t *testing.T) {
	result := longestCommonPrefix([]string{})
	if "" != result {
		t.Errorf("Result must be empty string for zero strings, actual is [%v]", result)
	}
}

func TestFindsFirstStringForSingleSet(t *testing.T) {
	strs := []string{"doc"}
	result := longestCommonPrefix(strs)
	if expected := "doc"; expected != result {
		t.Errorf("Result must be [%v] is first from set [%v], actual is [%v]", expected, strs, result)
	}
}

func TestFindsEmptyPrefixForDifferentStrings(t *testing.T) {
	strs := []string{"ac", "dc"}
	result := longestCommonPrefix(strs)
	if "" != result {
		t.Errorf("Result must be empty string for different strings [%v], actual is [%v]", strs, result)
	}
}

func TestFindsFullStringForTwoEqualsStrings(t *testing.T) {
	strs := []string{"doc", "doc"}
	result := longestCommonPrefix(strs)
	if expected := "doc"; expected != result {
		t.Errorf("Result must be [%v] for two equals strings [%v], actual is [%v]", expected, strs, result)
	}
}

func TestFindsFullStringForThreeEqualsString(t *testing.T) {
	strs := []string{"doc", "doc", "doc"}
	result := longestCommonPrefix(strs)
	if expected := "doc"; expected != result {
		t.Errorf("Result must be [%v] for three equals strings [%v], actual is [%v]", expected, strs, result)
	}
}

func TestFindsPrefixForTwoStringsWithEqualsStartsSubstringAndDifferentLength(t *testing.T) {
	strs := []string{"doc", "docs"}
	result := longestCommonPrefix(strs)
	if expected := "doc"; expected != result {
		t.Errorf("Result must be [%v] for string [%v] with equals start substring, actual is [%v]", expected, strs, result)
	}
}

func TestFindsPrefixForTwoStringsWithEqualsFirstChar(t *testing.T) {
	t.Skip()
	strs := []string{"aa", "ab"}
	result := longestCommonPrefix(strs)
	if expected := "a"; expected != result {
		t.Errorf("Result must be [%v] for first char equals string [%v], actual is [%v]", expected, strs, result)
	}
}
