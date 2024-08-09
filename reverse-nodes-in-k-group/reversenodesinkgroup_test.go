package reversenodesinkgroup

import "testing"

func TestReversesNoGroupForNilList(t *testing.T) {
	result := reverseKGroup(nil, 0)
	if nil != result {
		t.Errorf("Result must be nil for nil list, actual is [%v]", result)
	}
}

func TestReversesNoGroupForSingleNode(t *testing.T) {
	head, k := &ListNode{0, nil}, 1
	result := reverseKGroup(head, k)
	if nil == result || 0 != result.Val {
		t.Errorf("Result must be untouched head [%v] for k [%v], actual is [%v]", head, k, result)
	}
}
