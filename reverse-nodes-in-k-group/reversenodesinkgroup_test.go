package reversenodesinkgroup

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
	nums := []int{1, 2}
	head, k := makeListNodeFromNums(nums), 2
	result := reverseKGroup(head, k)
	if expectedNums, resultNums := [2]int{2, 1}, makeNumsFromListNode(result); 2 != len(resultNums) || expectedNums != [2]int(resultNums) {
		t.Errorf("Result must be nums [%v] for origin nums [%v], actual is [%v]", expectedNums, nums, resultNums)
	}
}

func TestReversesAllNodesForFourNodesList(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	head, k := makeListNodeFromNums(nums), 4
	result := reverseKGroup(head, k)
	if expectedNums, resultNums := [4]int{4, 3, 2, 1}, makeNumsFromListNode(result); 4 != len(resultNums) || expectedNums != [4]int(resultNums) {
		t.Errorf("Result must be nums [%v] for origin nums [%v], actual is [%v]", expectedNums, nums, resultNums)
	}
}
