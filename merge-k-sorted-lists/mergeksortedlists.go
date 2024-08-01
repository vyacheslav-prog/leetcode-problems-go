package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var mergedHead *ListNode
	for _, head := range lists {
		mergedHead = head
	}
	return mergedHead
}
