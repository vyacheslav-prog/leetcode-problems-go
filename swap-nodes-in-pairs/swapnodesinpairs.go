package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var leftHead, rightHead *ListNode
	mergedPreHead := &ListNode{0, nil}
	mergedTail := mergedPreHead
	for cursor := head; nil != cursor; cursor = cursor.Next {
		if nil == leftHead {
			leftHead = cursor
		} else if nil == rightHead {
			rightHead = cursor
		}
		if nil != rightHead {
			mergedTail.Next = &ListNode{rightHead.Val, &ListNode{leftHead.Val, nil}}
			mergedTail = mergedTail.Next.Next
			leftHead = nil
			rightHead = nil
		}
	}
	if nil != leftHead {
		mergedTail.Next = &ListNode{leftHead.Val, nil}
	}
	return mergedPreHead.Next
}
