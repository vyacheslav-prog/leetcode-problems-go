package palindromenumber

import "testing"

func TestPassesForZero(t *testing.T) {
	result := isPalindrome(0)
	if result != true {
		t.Errorf("Result must be [true] for [0], actual [%v]", result)
	}
}

func TestRejectsNegativeNumber(t *testing.T) {
	x := -1
	result := isPalindrome(x)
	if result != false {
		t.Errorf("Result must be [false] for negative [%v], actual is [%v]", x, result)
	}
}

func TestRejectsForTwoOtherDigits(t *testing.T) {
	x := 12
	result := isPalindrome(x)
	if result != false {
		t.Errorf("Result must be [false] for two other digits [%v], actual is [%v]", x, result)
	}
}

func TestPassesForTwoSameDigits(t *testing.T) {
	x := 11
	result := isPalindrome(x)
	if result != true {
		t.Errorf("Result must be [true] for two same digits [%v], actual is [%v]", x, result)
	}
}

func TestPassesForThreeSameDigits(t *testing.T) {
	x := 111
	result := isPalindrome(x)
	if result != true {
		t.Errorf("Result must be [true] for three same digits [%v], actual is [%v]", x, result)
	}
}

func TestRejectsForThreeOtherDigits(t *testing.T) {
	x := 123
	result := isPalindrome(x)
	if result != false {
		t.Errorf("Result must be [false] for three other digits [%v], actual is [%v]", x, result)
	}
}

func TestRejectsForFourDigitsWithSameFirstAndLast(t *testing.T) {
	x := 1231
	result := isPalindrome(x)
	if result != false {
		t.Errorf("Result must be [false] for four digist [%v] with same first and last, actual is [%v]", x, result)
	}
}

func TestPassesForFiveDigitsWithSymmetricalToMiddle(t *testing.T) {
	x := 12321
	result := isPalindrome(x)
	if result != true {
		t.Errorf("Result must be [true] for five digits [%v] with symmetrical, actual is [%v]", x, result)
	}
}

func TestRejectsNegativeSymmetricalToMiddleNumber(t *testing.T) {
	x := -101
	result := isPalindrome(x)
	if result != false {
		t.Errorf("Negative symmetrical to middle number [%v] it is [false] palindrome, actual is [%v]", x, result)
	}
}
