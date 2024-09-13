package dividetwointegers

import "testing"

func TestCalcsToOneForDivisorEqualsToDividend(t *testing.T) {
	dividend, divisor := 1, 1
	result := divide(dividend, divisor)
	if expected := 1; expected != result {
		t.Errorf("Result must be [%v] for equals dividend [%v] and divisor [%v], actual is [%v]", expected, dividend, divisor, result)
	}
}

func TestCalcsToDividentWhenDivisorIsOne(t *testing.T) {
	dividend := 2
	result := divide(dividend, 1)
	if dividend != result {
		t.Errorf("Result must be [%v], same as dividend for divisor is 1, actual is [%v]", dividend, result)
	}
}

func TestCalcsQuotientWhenDividendIsTwoDivisor(t *testing.T) {
	dividend, divisor := 6, 3
	result := divide(dividend, divisor)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for two-times dividend [%v] for divisor [%v], actual is [%v]", expected, dividend, divisor, result)
	}
}

func TestCalcsToZeroForDivisorGreaterThanDividend(t *testing.T) {
	dividend, divisor := 1, 2
	result := divide(dividend, divisor)
	if 0 != result {
		t.Errorf("Result must be 0 for divisor [%v] greater than dividend [%v], actual is [%v]", divisor, dividend, result)
	}
}

func TestCalcsToNegativeResultForNegativeDivisor(t *testing.T) {
	dividend, divisor := 2, -2
	result := divide(dividend, divisor)
	if expected := -1; expected != result {
		t.Errorf("Result must be [%v] for dividend [%v] and negative divisor [%v], actual is [%v]", expected, dividend, divisor, result)
	}
}
