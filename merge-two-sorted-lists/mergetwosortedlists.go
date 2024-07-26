package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var cursorHead **ListNode
	var mergedHead, mergedTail *ListNode
	for nil != list1 || nil != list2 {
		if nil == list2 {
			cursorHead = &list1
		} else if nil == list1 {
			cursorHead = &list2
		} else if list1.Val < list2.Val {
			cursorHead = &list1
		} else {
			cursorHead = &list2
		}
		if nil == mergedHead {
			mergedHead = &ListNode{(*cursorHead).Val, nil}
			mergedTail = mergedHead
		} else {
			mergedTail.Next = &ListNode{(*cursorHead).Val, nil}
			mergedTail = mergedTail.Next
		}
		*cursorHead = (*cursorHead).Next
	}
	return mergedHead
}
