package reversenodesinkgroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var buffer *ListNode
	var prevVal int
	for cursor := head; nil != cursor; cursor = cursor.Next {
		if nil == buffer {
			buffer = cursor
		} else {
			buffer.Val = cursor.Val
			buffer.Next.Val = prevVal
			buffer = nil
		}
		prevVal = cursor.Val
	}
	return head
}
