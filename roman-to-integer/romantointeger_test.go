package romantointeger

import "testing"

func TestRepresentsEmptyRomanNumberAsZero(t *testing.T) {
	result := romanToInt("")
	if 0 != result {
		t.Errorf("Result must be zero for empty roman number, actual is [%v]", result)
	}
}

func TestConvertsUnitFromRomanI(t *testing.T) {
	s := "I"
	result := romanToInt(s)
	if expected := 1; expected != result {
		t.Errorf("Result must be [%v] for a unit roman of [%v], actual is [%v]", expected, s, result)
	}
}

func TestConvertsTwoFromAppendedRomanII(t *testing.T) {
	s := "II"
	result := romanToInt(s)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for a two as appended symbols [%v], actual is [%v]", expected, s, result)
	}
}

func TestConvertsFourForSubstractiveRomanIV(t *testing.T) {
	s := "IV"
	result := romanToInt(s)
	if expected := 4; expected != result {
		t.Errorf("Result must be [%v] for a substractive roman [%v], actual is [%v]", expected, s, result)
	}
}

func TestConvertsFourteenForAppendenAndSubstractiveXIV(t *testing.T) {
	s := "XIV"
	result := romanToInt(s)
	if expected := 14; expected != result {
		t.Errorf("Result must be [%v] for an appended substractive roman [%v], actual is [%v]", expected, s, result)
	}
}

func TestConvertsRomanNumberWithManyAppendingAndSubstractings(t *testing.T) {
	s := "MCMXCIV"
	result := romanToInt(s)
	if expected := 1994; expected != result {
		t.Errorf("Result must be [%v] for a difference roman [%v], actual is [%v]", expected, s, result)
	}
}
