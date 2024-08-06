package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var mergedHead, mergedTail *ListNode
	var minimalHead **ListNode
	for {
		minimalHead = nil
		for index := 0; index != len(lists); index += 1 {
			if nil != lists[index] && (nil == minimalHead || lists[index].Val <= (*minimalHead).Val) {
				minimalHead = &lists[index]
			}
		}
		if nil == minimalHead {
			break
		}
		if nil == mergedHead {
			mergedHead = &ListNode{(*minimalHead).Val, nil}
			mergedTail = mergedHead
		} else {
			mergedTail.Next = &ListNode{(*minimalHead).Val, nil}
			mergedTail = mergedTail.Next
		}
		(*minimalHead) = (*minimalHead).Next
	}
	return mergedHead
}
