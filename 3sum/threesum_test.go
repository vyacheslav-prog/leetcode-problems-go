package threesum

import "testing"

func TestFindsZeroTripletsForEmptyNums(t *testing.T) {
	result := threeSum([]int{})
	if 0 != len(result) {
		t.Errorf("Result must be empty for an empty array nums, actual is [%v]", result)
	}
}

func TestFindsTripletsForAllNums(t *testing.T) {
	nums := []int{0, 0, 0}
	result := threeSum(nums)
	if expected := [3]int{0, 0, 0}; expected != [3]int(result[0]) {
		t.Errorf("Result must have [%v] for all nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestFindsZeroTripletsForNumsWithoutZeroSum(t *testing.T) {
	nums := []int{0, 0, 1}
	result := threeSum(nums)
	if 0 != len(result) {
		t.Errorf("Result must be empty for nums without a zero sum [%v], actual is [%v]", nums, result)
	}
}
