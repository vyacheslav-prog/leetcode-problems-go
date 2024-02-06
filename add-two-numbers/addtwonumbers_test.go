package addtwonumbers

import "testing"

func TestOutputsNilForNilLists(t *testing.T) {
	result := addTwoNumbers(nil, nil)
	if result != nil {
		t.Errorf("Result must be nil when list is nil, result: %v", result)
	}
}
