package mergetwosortedlists

import "testing"

func TestMergesToNilForTwoNilLists(t *testing.T) {
	result := mergeTwoLists(nil, nil)
	if nil != result {
		t.Errorf("Result must be nil for two nil lists, actual is [%v]", result)
	}
}

func TestEqualsToFirstNotNilListWhenSecondIsNil(t *testing.T) {
	first := &ListNode{0, nil}
	result := mergeTwoLists(first, nil)
	if nil == result {
		t.Errorf("Result must be [%v] for first [%v] and nil second, actual is [%v]", first, first, result)
	}
}

func TestEqualsToSecondNotNilListWhenFirstIsNil(t *testing.T) {
	second := &ListNode{0, nil}
	result := mergeTwoLists(nil, second)
	if nil == result {
		t.Errorf("Result must be [%v] for nil first and second [%v], actual is [%v]", second, second, result)
	}
}

func TestMergesTwoSingleValLists(t *testing.T) {
	first, second := &ListNode{0, nil}, &ListNode{1, nil}
	result := mergeTwoLists(first, second)
	if expected := (&ListNode{0, &ListNode{1, nil}}); nil == result || nil == result.Next {
		t.Errorf("Result must be [%v] for first [%v] and second [%v], actual is [%v]", expected, first, second, result)
	}
}
