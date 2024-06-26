package lettercombinationsofaphonenumber

import "testing"

func TestCombinatesForZeroDigits(t *testing.T) {
	result := letterCombinations("")
	if 0 != len(result) {
		t.Errorf("Result must be empty slice, actual is [%v]", result)
	}
}

func TestCombinatesForSingleDigit2(t *testing.T) {
	digits := "2"
	result := letterCombinations(digits)
	if expected := [3]string{"a", "b", "c"}; 3 != len(result) || expected != [3]string(result) {
		t.Errorf("Result must be [%v] for single digit [%v], actual is [%v]", expected, digits, result)
	}
}

func TestCombinatesForSingleDigit3(t *testing.T) {
	digits := "3"
	result := letterCombinations(digits)
	if expected := [3]string{"d", "e", "f"}; 3 != len(result) || expected != [3]string(result) {
		t.Errorf("Result must be [%v] for single digit [%v], actual is [%v]", expected, digits, result)
	}
}

func TestCombinatesForSingleDigit4(t *testing.T) {
	digits := "4"
	result := letterCombinations(digits)
	if expected := [3]string{"g", "h", "i"}; 3 != len(result) || expected != [3]string(result) {
		t.Errorf("Result must be [%v] for single digit [%v], actual is [%v]", expected, digits, result)
	}
}

func TestCombinatesForSingleDigit5(t *testing.T) {
	digits := "5"
	result := letterCombinations(digits)
	if expected := [3]string{"j", "k", "l"}; 3 != len(result) || expected != [3]string(result) {
		t.Errorf("Result must be [%v] for single digit [%v], actual is [%v]", expected, digits, result)
	}
}

func TestCombinatesForSingleDigit6(t *testing.T) {
	digits := "6"
	result := letterCombinations(digits)
	if expected := [3]string{"m", "n", "o"}; 3 != len(result) || expected != [3]string(result) {
		t.Errorf("Result must be [%v] for single digit [%v], actual is [%v]", expected, digits, result)
	}
}

func TestCombinatesForSingleDigit7(t *testing.T) {
	digits := "7"
	result := letterCombinations(digits)
	if expected := [4]string{"p", "q", "r", "s"}; 4 != len(result) || expected != [4]string(result) {
		t.Errorf("Result must be [%v] for single digit [%v], actual is [%v]", expected, digits, result)
	}
}

func TestCombinatesForSingleDigit8(t *testing.T) {
	digits := "8"
	result := letterCombinations(digits)
	if expected := [3]string{"t", "u", "v"}; 3 != len(result) || expected != [3]string(result) {
		t.Errorf("Result must be [%v] for single digit [%v], actual is [%v]", expected, digits, result)
	}
}

func TestCombinatesForSingleDigit9(t *testing.T) {
	digits := "9"
	result := letterCombinations(digits)
	if expected := [4]string{"w", "x", "y", "z"}; 4 != len(result) || expected != [4]string(result) {
		t.Errorf("Result must be [%v] for single digit [%v], actual is [%v]", expected, digits, result)
	}
}
