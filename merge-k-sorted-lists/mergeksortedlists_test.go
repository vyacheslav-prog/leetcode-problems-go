package mergeksortedlists

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

func TestMergesTwoListsWhenFirstTailLessThanHeadSecondList(t *testing.T) {
	firstNums, secondNums := []int{-3, -2, -1}, []int{1, 2, 3}
	lists := []*ListNode{makeListNodeFromNums(firstNums), makeListNodeFromNums(secondNums)}
	result := mergeKLists(lists)
	if expectedNums, resultNums := [6]int(append(firstNums, secondNums...)), makeNumsFromListNode(result); 6 != len(resultNums) || expectedNums != [6]int(resultNums) {
		t.Errorf("Result must be nums [%v] for two list [%v] and [%v], actual is [%v]", expectedNums, firstNums, secondNums, resultNums)
	}
}

func TestMergesThreeListsWithChainedTailToHead(t *testing.T) {
	firstNums, secondNums, thirdNums := []int{-1000, -100}, []int{-10, 10}, []int{100, 1000}
	lists := []*ListNode{makeListNodeFromNums(firstNums), makeListNodeFromNums(secondNums), makeListNodeFromNums(thirdNums)}
	result := mergeKLists(lists)
	if expectedNums, resultNums := [6]int(append(firstNums, append(secondNums, thirdNums...)...)), makeNumsFromListNode(result); 6 != len(resultNums) || expectedNums != [6]int(resultNums) {
		t.Errorf("Result must be nums [%v] for chained lists [%v], [%v] and [%v], actual nums is [%v]", expectedNums, firstNums, secondNums, thirdNums, resultNums)
	}
}

func TestMergesTwoListsWhenSecondTailLessThanHeadFirstList(t *testing.T) {
	firstNums, secondNums := []int{1, 2}, []int{-1, 0}
	lists := []*ListNode{makeListNodeFromNums(firstNums), makeListNodeFromNums(secondNums)}
	result := mergeKLists(lists)
	if expectedNums, resultNums := [4]int{-1, 0, 1, 2}, makeNumsFromListNode(result); 4 != len(resultNums) || expectedNums != [4]int(resultNums) {
		t.Errorf("Result must be nums [%v] for tail-head list nums [%v] and [%v], actual nums is [%v]", expectedNums, firstNums, secondNums, resultNums)
	}
}

func TestMergesThreeListTailToHead(t *testing.T) {
	firstNums, secondNums, thirdNums := []int{2, 3}, []int{-2, -1}, []int{0, 1}
	lists := []*ListNode{makeListNodeFromNums(firstNums), makeListNodeFromNums(secondNums), makeListNodeFromNums(thirdNums)}
	result := mergeKLists(lists)
	if expectedNums, resultNums := [6]int{-2, -1, 0, 1, 2, 3}, makeNumsFromListNode(result); 6 != len(resultNums) || expectedNums != [6]int(resultNums) {
		t.Errorf("Result must be nums [%v] for tail-head list nums [%v], [%v] and [%v], actual nums is [%v]", expectedNums, firstNums, secondNums, thirdNums, resultNums)
	}
}
