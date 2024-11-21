package searchinrotatedsortedarray

import "testing"

func TestSearchesNoIndexForEmptyArray(t *testing.T) {
	result := search(nil, 0)
	if -1 != result {
		t.Errorf("Result must be -1 for empty array, actual is [%v]", result)
	}
}
