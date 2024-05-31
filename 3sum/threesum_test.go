package threesum

import "testing"

func TestFindsZeroTripletsForEmptyNums(t *testing.T) {
	result := threeSum([]int{})
	if 0 != len(result) {
		t.Errorf("Result must be empty for an empty array nums, actual is [%v]", result)
	}
}
