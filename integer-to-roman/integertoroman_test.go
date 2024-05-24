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

func TestConvertsWithAppendingOtherSymbols(t *testing.T) {
	num := 6
	result := intToRoman(num)
	if expected := "VI"; expected != result {
		t.Errorf("Result must be appending [%v] for number [%v], actual is [%v]", expected, num, result)
	}
}

func TestConvertsWithSubstractingOtherSymbol(t *testing.T) {
	num := 4
	result := intToRoman(num)
	if expected := "IV"; expected != result {
		t.Errorf("Result must be substractive [%v] for number [%v], actual is [%v]", expected, num, result)
	}
}

func TestConvertsWithAppendingAndSubstractingOtherSymbol(t *testing.T) {
	num := 19
	result := intToRoman(num)
	if expected := "XIX"; result != expected {
		t.Errorf("Result must be appended and substracted [%v] for number [%v], actual is [%v]", expected, num, result)
	}
}

func TestConvertsWithAppendingAndSubstractingOtherSymbolForThousands(t *testing.T) {
	num := 1900
	result := intToRoman(num)
	if expected := "MCM"; result != expected {
		t.Errorf("Result must be appended and substracted [%v] for number [%v], actual is [%v]", expected, num, result)
	}
}
