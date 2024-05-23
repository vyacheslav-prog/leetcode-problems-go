package integertoroman

import "testing"

func TestRepresentsEmptyStringForZero(t *testing.T) {
	result := intToRoman(0)
	if "" != result {
		t.Errorf("Result must be empty string for number [%v], actual is [%v]", 0, result)
	}
}

func TestConvertsUnitAsI(t *testing.T) {
	num := 1
	result := intToRoman(num)
	if expected := "I"; expected != result {
		t.Errorf("Result must be [%v] for number [%v], actual is [%v]", expected, num, result)
	}
}

func TestConvertsFiveAsV(t *testing.T) {
	num := 5
	result := intToRoman(num)
	if expected := "V"; expected != result {
		t.Errorf("Result must be [%v] for number [%v], actual is [%v]", expected, num, result)
	}
}

func TestConvertsTwoAsRepeatedUnitSymbol(t *testing.T) {
	num := 2
	result := intToRoman(num)
	if expected := "II"; expected != result {
		t.Errorf("Result must be [%v] for number [%v], actual is [%v]", expected, num, result)
	}
}
