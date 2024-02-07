package addtwonumbers

import "testing"

func TestOutputsNilForNilLists(t *testing.T) {
	result := addTwoNumbers(nil, nil)
	if result != nil {
		t.Errorf("Result must be nil when list is nil, result: %v", result)
	}
}

func TestSumsEmptyListsToEmptyList(t *testing.T) {
	l1 := &ListNode{}
	l2 := &ListNode{}
	result := addTwoNumbers(l1, l2)
	if result.Val != 0 || result.Next != nil {
		t.Errorf("Result must be zero when lists is empty, result: %v", result)
	}
}

func TestSumsSingletonLists(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 1}
	result := addTwoNumbers(l1, l2)
	if result.Val != 2 || result.Next != nil {
		t.Errorf("Result must be a sum two nums without added nums, result: %v", result)
	}
}

func TestSumsSingletonListsWithOverflow(t *testing.T) {
	l1 := &ListNode{Val: 5}
	l2 := &ListNode{Val: 5}
	result := addTwoNumbers(l1, l2)
	if result.Val != 0 || result.Next.Val != 1 {
		t.Errorf("Result must be overflowed from first value, result: %v", result)
	}
}

func TestSumsSymmetricMultipleLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 1}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 1}}
	result := addTwoNumbers(l1, l2)
	if result.Val != 2 || result.Next.Val != 2 {
		t.Errorf("Must sum the symmetric multiple lists without overflow, result: %v", result)
	}
}

func TestSumsSymmetricMultipleListsWithOverflow(t *testing.T) {
	l1 := &ListNode{Val: 5, Next: &ListNode{Val: 1}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 1}}
	result := addTwoNumbers(l1, l2)
	if result.Val != 0 || result.Next.Val != 3 {
		t.Errorf("Must sum the symmetric lists with overflow, first: %v, second: %v", result.Val, result.Next.Val)
	}
}

func TestSumsUnsymmetricLists(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 0, Next: &ListNode{Val: 1}}
	result := addTwoNumbers(l1, l2)
	if result.Val != 1 || result.Next.Val != 1 {
		t.Errorf("Must sum the unsymmetric lists without overflow, first: %v, second: %v", result.Val, result.Next.Val)
	}
}

func TestSumsUnsymmetricListsWithOverflow(t *testing.T) {
	l1 := &ListNode{Val: 5}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 1}}
	result := addTwoNumbers(l1, l2)
	if result.Val != 0 || result.Next.Val != 2 {
		t.Errorf("Must sum the unsymmetric lists with overflow, first: %v, second: %v", result.Val, result.Next.Val)
	}
}

func TestSumsUnsymmetricListsWithManyOverflow(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}
	result := addTwoNumbers(l1, l2)
	if result.Val != 8 || result.Next.Val != 9 || result.Next.Next.Val != 0 || result.Next.Next.Next.Val != 1 {
		t.Errorf(
			"Must sum the unsymmetric lists with many overflow, nums: %v %v %v %v",
			result.Val,
			result.Next.Val,
			result.Next.Next.Val,
			result.Next.Next.Next.Val,
		)
	}
}
