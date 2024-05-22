package integertoroman

import "testing"

func TestRepresentsEmptyStringForZero(t *testing.T) {
	result := intToRoman(0)
	if "" != result {
		t.Errorf("Result must be empty string for number [%v], actual is [%v]", 0, result)
	}
}
