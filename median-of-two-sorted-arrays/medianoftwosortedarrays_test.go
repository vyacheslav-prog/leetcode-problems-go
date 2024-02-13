package medianoftwosortedarrays

import "testing"

func TestFindsZeroWhenNumsIsEmpty(t *testing.T) {
	result := findMedianSortedArrays([]int{}, []int{})
	if result != 0 {
		t.Errorf("Result must be a zero when nums is empty, result: %v", result)
	}
}

func TestFindsMedianEqualsItemForSingletonNums(t *testing.T) {
	nums := []int{1}
	result1, result2 := findMedianSortedArrays(nums, []int{}), findMedianSortedArrays([]int{}, nums)
	if result1 != 1 || result1 != result2 {
		t.Errorf("Result must be 1 for singleton %v, first result: %v, second result: %v", nums, result1, result2)
	}
}

func TestFindsMedianForTwoSingletonSortedNums(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{2}
	result1, result2 := findMedianSortedArrays(nums1, nums2), findMedianSortedArrays(nums2, nums1)
	if result1 != 1.5 || result1 != result2 {
		t.Errorf("Result must be 1.5 for two %v %v, first result: %v, second result: %v", nums1, nums2, result1, result2)
	}
}

func TestFindsMedianForOddMultipleNumsAndEmptyNums(t *testing.T) {
	nums := []int{1, 2, 3}
	result1, result2 := findMedianSortedArrays(nums, []int{}), findMedianSortedArrays([]int{}, nums)
	if result1 != 2.0 || result1 != result2 {
		t.Errorf("Result must be 2.0 for odd %v and empty other nums, first result: %v, second result: %v", nums, result1, result2)
	}
}

func TestFindsMedianForEvenMultipleNumsAndEmptyNums(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	result1, result2 := findMedianSortedArrays(nums, []int{}), findMedianSortedArrays([]int{}, nums)
	if result1 != 2.5 || result1 != result2 {
		t.Errorf("Result must be 2.5 for even %v and empty other nums, first result: %v, second result: %v", nums, result1, result2)
	}
}

func TestFindsMedianForEvenMultipleNums(t *testing.T) {
	nums1 := []int{1,2}
	nums2 := []int{3,4}
	result1, result2 := findMedianSortedArrays(nums1, nums2), findMedianSortedArrays(nums2, nums1)
	if result1 != 2.5 || result1 != result2 {
		t.Errorf("Result must be 2.5 for even %v and %v, first result: %v, second result: %v", nums1, nums2, result1, result2)
	}
}
