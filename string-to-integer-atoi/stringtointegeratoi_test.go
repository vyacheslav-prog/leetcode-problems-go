package stringtointegeratoi

import "testing"

func TestConvertsEmptyStringToZero(t *testing.T) {
	result := myAtoi("")
	if result != 0 {
		t.Errorf("Empty string must be converted to zero, actual [%v]", result)
	}
}

func TestConvertsOneDigitString(t *testing.T) {
	s := "1"
	result := myAtoi(s)
	if result != 1 {
		t.Errorf("String with one digit [%v] must returns [1], actual [%v]", s, result)
	}
}

func TestConvertsOneDigitWithWhitespace(t *testing.T) {
	s := " 2"
	result := myAtoi(s)
	if result != 2 {
		t.Errorf("One digit string with leading whitespace [%v] must converting to [2], actual [%v]", s, result)
	}
}

func TestConvertsNegativeOneDigitString(t *testing.T) {
	s := "-1"
	result := myAtoi(s)
	if result != -1 {
		t.Errorf("String with leading minus [%v] must converting to [-1], actual [%v]", s, result)
	}
}

func TestConvertsTwoDigitsString(t *testing.T) {
	s := "12"
	result := myAtoi(s)
	if result != 12 {
		t.Errorf("String with two digits [%v] must converting to [12], actual [%v]", s, result)
	}
}

func TestConvertsWithClampToMaxNumber(t *testing.T) {
	s := "2147483648"
	result := myAtoi(s)
	if result != 2147483647 {
		t.Errorf("String with overflow [%v] must be clamped to [2147483647], actual [%v]", s, result)
	}
}
