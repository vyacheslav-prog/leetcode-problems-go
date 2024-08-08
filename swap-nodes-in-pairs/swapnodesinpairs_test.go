package swapnodesinpairs

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

func TestSwapsTwoNodeList(t *testing.T) {
	nums := []int{0, 1}
	result := swapPairs(makeListNodeFromNums(nums))
	if expectedNums, resultNums := [2]int{1, 0}, makeNumsFromListNode(result); 2 != len(resultNums) || expectedNums != [2]int(resultNums) {
		t.Errorf("Result must have nums [%v] for list with nums [%v], actual is [%v]", expectedNums, nums, resultNums)
	}
}

func TestSwapsTwoHeadNodesIntoThreeNodeList(t *testing.T) {
	nums := []int{-2, -1, -3}
	result := swapPairs(makeListNodeFromNums(nums))
	if expectedNums, resultNums := [3]int{-1, -2, -3}, makeNumsFromListNode(result); 3 != len(resultNums) || expectedNums != [3]int(resultNums) {
		t.Errorf("Result must have nums [%v] for list with nums [%v], actual is [%v]", expectedNums, nums, resultNums)
	}
}

func TestSwapsTwoPairNodesIntoFourNodeList(t *testing.T) {
	nums := []int{1, 0, 3, 2}
	result := swapPairs(makeListNodeFromNums(nums))
	if expectedNums, resultNums := [4]int{0, 1, 2, 3}, makeNumsFromListNode(result); 4 != len(resultNums) || expectedNums != [4]int(resultNums) {
		t.Errorf("Result must have nums [%v] for list with nums [%v], actual is [%v]", expectedNums, nums, resultNums)
	}
}
