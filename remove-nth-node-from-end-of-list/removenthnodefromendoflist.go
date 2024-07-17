package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var tailOffset int
	preTail, tail := head, head
	for cursor := head; nil != cursor; cursor = cursor.Next {
		if tailOffset < n {
			tailOffset += 1
		} else {
			preTail = tail
			tail = tail.Next
		}
	}
	if nil == head || tailOffset < n {
		head = tail
	} else if head == tail {
		head = head.Next
	} else {
		preTail.Next = tail.Next
	}
	return head
}
