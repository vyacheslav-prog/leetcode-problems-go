package removenthnodefromendoflist

import "testing"

func TestTransformsNilNodeListToNil(t *testing.T) {
	result := removeNthFromEnd(nil, 0)
	if nil != result {
		t.Errorf("Result must be nil for a nil list, actual is [%v]", result)
	}
}

func TestUntouchesNodeListWhenNthNodeIsOverBounds(t *testing.T) {
	head, n := &ListNode{0, nil}, 2
	result := removeNthFromEnd(head, n)
	if head != result {
		t.Errorf("Result must be equals for head [%v] for over bounds nth [%v], actual is [%v]", head, n, result)
	}
}

func TestTransformsToNilForSingletonListAnd1thNode(t *testing.T) {
	head, n := &ListNode{0, nil}, 1
	result := removeNthFromEnd(head, n)
	if nil != result {
		t.Errorf("Result must be nil for singleton list [%v] and nth node [%v], actual is [%v]", head, n, result)
	}
}
