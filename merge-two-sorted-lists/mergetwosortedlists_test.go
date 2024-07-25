package mergetwosortedlists

import "testing"

func makeListNodeFromNums(nums []int) *ListNode {
	var cursor, head *ListNode
	for index, value := range nums {
		if 0 == index {
			head = &ListNode{value, nil}
			cursor = head
		} else {
			cursor.Next = &ListNode{value, nil}
			cursor = cursor.Next
		}
	}
	return head
}

func makeNumsFromListNode(head *ListNode) []int {
	var result []int
	for nil != head {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

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

func TestMergesTwoSingleValListsWhenGreaterValIntoFirstList(t *testing.T) {
	first, second := &ListNode{1, nil}, &ListNode{0, nil}
	result := mergeTwoLists(first, second)
	if expectedHeadVal := 0; nil == result || expectedHeadVal != result.Val {
		t.Errorf("Result must have first val [%v] for lists [%v] and [%v], actual list is [%v]", expectedHeadVal, first, second, result)
	}
}

func TestMergesTwoUnorderedForOtherList(t *testing.T) {
	firstVals, secondVals := []int{1, 3}, []int{0, 4}
	first, second := makeListNodeFromNums(firstVals), makeListNodeFromNums(secondVals)
	result := mergeTwoLists(first, second)
	if expectedVals, resultVals := [4]int{0, 1, 3, 4}, makeNumsFromListNode(result); 4 != len(resultVals) || expectedVals != [4]int(resultVals) {
		t.Errorf("Result must have vals [%v] for vals [%v] and [%v], actual vals is [%v]", expectedVals, firstVals, secondVals, resultVals)
	}
}
