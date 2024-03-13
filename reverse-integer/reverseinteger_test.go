package reverseinteger

import "testing"

func TestReversesZero(t *testing.T) {
	result := reverse(0)
	if result != 0 {
		t.Error("Result must be [0] for number [0]")
	}
}

func TestReversesSingleDigit(t *testing.T) {
	x := 1
	result := reverse(x)
	if result != 1 {
		t.Errorf("Result must be [1] for single digit [%v], actual [%v]", x, result)
	}
}

func TestReversesTwoDigitsNumber(t *testing.T) {
	x := 12
	result := reverse(x)
	if result != 21 {
		t.Errorf("Result must be [21] for two digits [%v], actual [%v]", x, result)
	}
}

func TestReversesSingleNegativeDigit(t *testing.T) {
	x := -1
	result := reverse(x)
	if result != -1 {
		t.Errorf("Result must be [-1] for single negative [%v], actual [%v]", x, result)
	}
}

func TestReversesTwoDigitsWithLeadingZero(t *testing.T) {
	x := 10
	result := reverse(x)
	if result != 1 {
		t.Errorf("Result must be [1] for number with leading zero [%v], actual [%v]", x, result)
	}
}

func TestReversesTwoDigitsNumberWhenNegative(t *testing.T) {
	x := -12
	result := reverse(x)
	if result != -21 {
		t.Errorf("Result must be [-21] for two negitve digits [%v], actual [%v]", x, result)
	}
}

func TestReversesThreeDigits(t *testing.T) {
	x := 123
	result := reverse(x)
	if result != 321 {
		t.Errorf("Result must be [321] for three digits [%v], actual[%v]", x, result)
	}
}

func TestReversesThreeDigitsWhenSignIsNegative(t *testing.T) {
	x := -321
	result := reverse(x)
	if result != -123 {
		t.Errorf("Result must be [-123] for three digits with negative sign [%v], actual [%v]", x, result)
	}
}

func TestReturnsZeroWhenNumberWithOverflow(t *testing.T) {
	x := 2147483648 // 2^31
	result := reverse(x)
	if result != 0 {
		t.Errorf("Result must be [0] when a number reversing for [%v] is overflowed, actual [%v]", x, result)
	}
}

func TestReturnsZeroWhenNumberIsNegativeOverflow(t *testing.T) {
	x := -2147483649
	result := reverse(x)
	if result != 0 {
		t.Errorf("Result must be [0] when a negative number [%v] is overflowed reversing, actual [%v]", x, result)
	}
}

func TestReturnsZeroWhenReversingIsOverflowAndStartsFromOne(t *testing.T) {
	x := 1534236469
	result := reverse(x)
	if result != 0 {
		t.Errorf("Result must be [0] for an overflowed reversing [%v] with [1] on start, actual [%v]", x, result)
	}
}

func TestReturnsZeroWhenReversingIsOverflowAndStartsFromOneForNegativeNumber(t *testing.T) {
	x := -1534236469
	result := reverse(x)
	if result != 0 {
		t.Errorf("Result must be [0] for an negative overflowed reversing [%v] with [1] on start, actual [%v]", x, result)
	}
}

func TestReversesPositiveLargeNumberWithZeroIntoMiddle(t *testing.T) {
	x := 24077
	result := reverse(x)
	if result != 77042 {
		t.Errorf("Result must be [77042] for a positive large number [%v], actual [%v]", x, result)
	}
}

func TestReversesNegativeMaxDigitsNumberWithoutOverflow(t *testing.T) {
	x := -2147483412
	result := reverse(x)
	if result != -2143847412 {
		t.Errorf("Result must be [-2143847412] for negative max digits [%v], actual [%v]", x, result)
	}
}

func TestReversesPositivePreMaxDigitsNumberGreatThanPositiveMaxWithoutRightDigit(t *testing.T) {
	x := 463847412
	result := reverse(x)
	if result != 214748364 {
		t.Errorf("Result must be [214748364] for pre max digits [%v], actual [%v]", x, result)
	}
}

func TestReturnsZeroWhenReversingIsOverflowedAndFirstDigitIsOne(t *testing.T) {
	x := 1563847412
	result := reverse(x)
	if result != 0 {
		t.Errorf("Result must be [0] for an overflowed reversing [%v] with one on start, actual [%v]", x, result)
	}
}
