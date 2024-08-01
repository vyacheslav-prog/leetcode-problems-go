package mergeksortedlists

import "testing"

func TestMergesToNilForEmptyLists(t *testing.T) {
	result := mergeKLists(nil)
	if nil != result {
		t.Errorf("Result must be nil for empty lists, actual is [%v]", result)
	}
}

func TestMergesToFirstListForSingletonList(t *testing.T) {
	lists := []*ListNode{&ListNode{0, nil}}
	result := mergeKLists(lists)
	if nil == result || 0 != result.Val {
		t.Errorf("Result must be first head for singleton list [%v], actual is [%v]", lists, result)
	}
}
