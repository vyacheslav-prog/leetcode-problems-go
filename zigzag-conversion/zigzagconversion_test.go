package zigzagconversion

import "testing"

func TestConvertsEmptyString(t *testing.T) {
	result := convert("", 0)
	if result != "" {
		t.Errorf("Result must be empty for empty string, actual [%v]", result)
	}
}

func TestConvertsOneCharString(t *testing.T) {
	numRows := 1
	s := "A"
	result := convert(s, numRows)
	if result != "A" {
		t.Errorf("Result must be [A] for string [%v] with num rows [%v], actual [%v]", s, numRows, result)
	}
}
