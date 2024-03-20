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
