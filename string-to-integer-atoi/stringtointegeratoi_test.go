package stringtointegeratoi

import "testing"

func TestConvertsEmptyStringToZero(t *testing.T) {
	result := myAtoi("")
	if result != 0 {
		t.Errorf("Empty string must be converted to zero, actual [%v]", result)
	}
}
