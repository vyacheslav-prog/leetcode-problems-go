package swapnodesinpairs

import "testing"

func TestSwapsNotForNilHead(t *testing.T) {
	result := swapPairs(nil)
	if nil != result {
		t.Errorf("Result must be nil for nil head, actual is [%v]", result)
	}
}

func TestSwapsOneNodeList(t *testing.T) {
	head := &ListNode{0, nil}
	result := swapPairs(head)
	if nil == result || 0 != result.Val {
		t.Errorf("Result must be same head [%v] after swap, actual is [%v]", head, result)
	}
}
