package addtwonumbers

import "testing"

func TestOutputsNilForNilLists(t *testing.T) {
	result := addTwoNumbers(nil, nil)
	if result != nil {
		t.Errorf("Result must be nil when list is nil, result: %v", result)
	}
}

func TestSumsEmptyListsToEmptyList(t *testing.T) {
	l1 := &ListNode{}
	l2 := &ListNode{}
	result := addTwoNumbers(l1, l2)
	if result.Val != 0 || result.Next != nil {
		t.Errorf("Results must be zero when lists is empty, result %v", result)
	}
}
