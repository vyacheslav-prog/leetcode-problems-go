package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var counter int
	cursor := head
	for nil != cursor {
		counter += 1
		if n == counter {
			if 1 == counter && nil == cursor.Next {
				head = nil
			}
			cursor.Next = nil
		}
		cursor = cursor.Next
	}
	return head
}
