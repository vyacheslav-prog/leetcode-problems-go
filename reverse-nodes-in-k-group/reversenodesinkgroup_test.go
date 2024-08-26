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

func TestReversesAllNodesForTwoNodesList(t *testing.T) {
	head, k := &ListNode{1, &ListNode{2, nil}}, 2
	result := reverseKGroup(head, k)
	if expectedHeadVal, expectedTailVal := 2, 1; nil == result || expectedHeadVal != result.Val || nil == result.Next || expectedTailVal != result.Next.Val {
		t.Errorf("Result must have reversed Val for head [%v] and tail [%v] for list [%v] and group [%v], actual head is [%v]", expectedHeadVal, expectedTailVal, head, k, result)
	}
}
