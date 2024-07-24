package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var mergedList *ListNode = list2
	if nil != list1 {
		mergedList = list1
		mergedList.Next = list2
	}
	return mergedList
}
