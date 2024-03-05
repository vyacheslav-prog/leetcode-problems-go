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

func TestConvertsManyCharsWhenRowIsSingle(t *testing.T) {
	numRows := 1
	s := "ABC"
	result := convert(s, numRows)
	if result != "ABC" {
		t.Errorf("Result must be [ABC] for string [%v] when a row is single, actual [%v]", s, result)
	}
}

func TestConvertsToTwoRows(t *testing.T) {
	numRows := 2
	s := "ABC"
	result := convert(s, numRows)
	if result != "ACB" {
		t.Errorf("Result must be [ACB] for string [%v] when a rows is twice, actual [%v]", s, result)
	}
}

func TestConvertsToTwoRowsWhenLengthIsMore(t *testing.T) {
	numRows := 2
	s := "ABCDEF"
	result := convert(s, numRows)
	if result != "ACEBDF" {
		t.Errorf("Result must be [ACEBDF] for string [%v] when a rows is twice, actual [%v]", s, result)
	}
}

func TestConvertsToThreeRowsWhenLengthIsThree(t *testing.T) {
	numRows := 3
	s := "123"
	result := convert(s, numRows)
	if result != "123" {
		t.Errorf("Result must be [123] for string [%v] when a rows is triple, actual [%v]", s, result)
	}
}

func TestConvertsToThreeRowsWhenColumnsIsMore(t *testing.T) {
	numRows := 3
	s := "PAYPALISHIRING"
	result := convert(s, numRows)
	if result != "PAHNAPLSIIGYIR" {
		t.Errorf("Result must be [PAHNAPLSIIGYIR] for string [%v] when a columns is more, actual [%v]", s, result)
	}
}

func TestConvertsToFourRowsWhenColumnsIsRepeated(t *testing.T) {
	numRows := 4
	s := "PAYPALISHIRING"
	result := convert(s, numRows)
	if result != "PINALSIGYAHRPI" {
		t.Errorf("Result must be [PINALSIGYAHRPI] for [%v] when rows is four and columns is repeated, actual [%v]", s, result)
	}
}
