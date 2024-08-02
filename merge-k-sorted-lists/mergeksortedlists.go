package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var mergedHead, mergedTail *ListNode
	for _, head := range lists {
		if nil == mergedTail {
			mergedHead = head
		}
		for cursor := mergedHead; nil != cursor; cursor = cursor.Next {
			mergedTail = cursor
		}
		if mergedHead != head {
			mergedTail.Next = head
		}
	}
	return mergedHead
}
