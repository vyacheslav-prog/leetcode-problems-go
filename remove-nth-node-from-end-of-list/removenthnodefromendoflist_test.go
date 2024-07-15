package removenthnodefromendoflist

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

func TestTransformsNilNodeListToNil(t *testing.T) {
	result := removeNthFromEnd(nil, 0)
	if nil != result {
		t.Errorf("Result must be nil for a nil list, actual is [%v]", result)
	}
}

func TestUntouchesNodeListWhenNthNodeIsOverBounds(t *testing.T) {
	head, n := makeListNodeFromNums([]int{0}), 2
	result := removeNthFromEnd(head, n)
	if head != result {
		t.Errorf("Result must be equals for head [%v] for over bounds nth [%v], actual is [%v]", head, n, result)
	}
}

func TestTransformsToNilForSingletonListAnd1thNode(t *testing.T) {
	head, n := makeListNodeFromNums([]int{0}), 1
	result := removeNthFromEnd(head, n)
	if nil != result {
		t.Errorf("Result must be nil for singleton list [%v] and nth node [%v], actual is [%v]", head, n, result)
	}
}

func TestRemovesLastNodeForTwoNodes(t *testing.T) {
	head, n := makeListNodeFromNums([]int{0, 1}), 1
	result := removeNthFromEnd(head, n)
	if nil == result || nil != result.Next {
		t.Errorf("Result must have head node only for list [%v] and nth node [%v], actual is [%v]", head, n, result)
	}
}

func TestRemovesLastNodeForThreeNodes(t *testing.T) {
	nums := []int{0, 1, 2}
	head, n := makeListNodeFromNums(nums), 1
	result := removeNthFromEnd(head, n)
	if expected, resultNums := [2]int{0, 1}, makeNumsFromListNode(result); 2 != len(resultNums) || expected != [2]int(resultNums) {
		t.Errorf("Result must nums [%v] for list [%v] and nth node [%v], actual is [%v]", expected, nums, n, resultNums)
	}
}
