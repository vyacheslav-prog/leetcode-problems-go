package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return addTwoNumbers(&ListNode{Val: 0}, l2)
	} else if l2 == nil {
		return addTwoNumbers(l1, &ListNode{Val: 0})
	}
	sum := l1.Val + l2.Val
	if sum > 9 {
		tens := sum / 10
		units := sum % 10
		if l1.Next != nil {
			return &ListNode{
				Val: units,
				Next: addTwoNumbers(
					&ListNode{
						Val:  l1.Next.Val + tens,
						Next: l1.Next.Next,
					},
					l2.Next,
				),
			}
		}
		if l2.Next != nil {
			return &ListNode{
				Val: units,
				Next: addTwoNumbers(
					l1.Next,
					&ListNode{
						Val:  l2.Next.Val + tens,
						Next: l2.Next.Next,
					},
				),
			}
		}
		return &ListNode{
			Val:  units,
			Next: &ListNode{Val: tens},
		}
	}
	if l1.Next != nil || l2.Next != nil {
		return &ListNode{
			Val:  sum,
			Next: addTwoNumbers(l1.Next, l2.Next),
		}
	}
	return &ListNode{Val: sum}
}
