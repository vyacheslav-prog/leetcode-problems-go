package reversenodesinkgroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseNodesIntoGroup(head *ListNode, offset int) {
	var swapCursor *ListNode
	var prevVal int
	for 0 != offset {
		for swapCounter := 0; offset != (swapCounter - 1); swapCounter += 1 {
			if nil == swapCursor {
				swapCursor = head
			} else {
				swapCursor.Val = swapCursor.Next.Val
				swapCursor.Next.Val = prevVal
				swapCursor = swapCursor.Next
			}
			prevVal = swapCursor.Val
		}
		swapCursor = nil
		offset -= 1
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var groupHead *ListNode
	var groupCounter int
	for cursor := head; nil != cursor; cursor = cursor.Next {
		if (k - 1) == groupCounter {
			reverseNodesIntoGroup(groupHead, groupCounter)
			groupCounter = 0
		} else {
			if 0 == groupCounter {
				groupHead = cursor
			}
			groupCounter += 1
		}
	}
	return head
}
