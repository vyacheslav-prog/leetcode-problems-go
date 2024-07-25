package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var cursorHead **ListNode
	var mergedList *ListNode
	for nil != list1 || nil != list2 {
		if nil != list1 && nil != list2 {
			if list1.Val < list2.Val {
				cursorHead = &list1
			} else {
				cursorHead = &list2
			}
		} else if nil == list2 {
			cursorHead = &list1
		} else {
			cursorHead = &list2
		}
		if nil == mergedList {
			mergedList = &ListNode{(*cursorHead).Val, nil}
		} else {
			mergedList.Next = &ListNode{(*cursorHead).Val, nil}
		}
		*cursorHead = (*cursorHead).Next
	}
	return mergedList
}
