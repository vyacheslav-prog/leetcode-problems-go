package findtheindexofthefirstoccurrenceinastring

import "testing"

func TestFindsNoIndexForEmptyHaystackAndNeedle(t *testing.T) {
	result := strStr("", "")
	if -1 != result {
		t.Errorf("Result must be -1 for empty haystack and needle, actual is [%v]", result)
	}
}

func TestFindsIndexForSingleCharAndEqualsStrings(t *testing.T) {
	haystack, needle := "a", "a"
	result := strStr(haystack, needle)
	if expected := 0; expected != result {
		t.Errorf("Result must be [%v] for single char haystack [%v] and equal needle [%v], actual is [%v]", expected, haystack, needle, result)
	}
}

func TestFindsNoIndexForEmptyNeedleAndNotEmptyHaystack(t *testing.T) {
	haystack, needle := "bc", ""
	result := strStr(haystack, needle)
	if -1 != result {
		t.Errorf("Result must be -1 for empty needle and not empty haystack [%v], actual is [%v]", haystack, result)
	}
}

func TestFindsLastCharIndexWhenNeedleIsLastChar(t *testing.T) {
	haystack, needle := "cat", "t"
	result := strStr(haystack, needle)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for last char needle [%v] into haystack [%v], actual is [%v]", expected, needle, haystack, result)
	}
}

func TestFindsNoIndexForNotPresenceSingletonCharNeedleIntoHaystack(t *testing.T) {
	haystack, needle := "dog", "z"
	result := strStr(haystack, needle)
	if -1 != result {
		t.Errorf("Result must be -1 for not presence needle [%v] into haystack [%v], actual is [%v]", needle, haystack, result)
	}
}

func TestFindsForSingletonCharNeedleIntoHaystackWhenPlaceIsMiddle(t *testing.T) {
	haystack, needle := "fox", "o"
	result := strStr(haystack, needle)
	if expected := 1; expected != result {
		t.Errorf("Result must be [%v] for needle [%v] into haystack [%v], actual is [%v]", expected, needle, haystack, result)
	}
}
