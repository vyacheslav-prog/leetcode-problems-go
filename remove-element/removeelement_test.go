package removeelement

import "testing"

func TestRemovesZeroNumsForEmptyNums(t *testing.T) {
	var nums []int
	result := removeElement(nums, 0)
	if 0 != result || nil != nums {
		t.Errorf("Result must be zero for nil nums [%v], actual is [%v]", nums, result)
	}
}

func TestRemovesZeroNumsForMissedVal(t *testing.T) {
	nums, val := []int{0, 1, 2}, 3
	result := removeElement(nums, val)
	if expectedNums, expectedVal := [3]int{0, 1, 2}, 3; expectedVal != result || 3 != len(nums) || expectedNums != [3]int(nums) {
		t.Errorf("Result must be [%v] for nums [%v] and missed val [%v], actual is [%v] for nums [%v]", expectedVal, expectedNums, val, result, nums)
	}
}

func TestRemovesSingletonValForStartIntoNums(t *testing.T) {
	nums, val := []int{5, 1, 2}, 5
	result := removeElement(nums, val)
	if expectedNums, expectedVal := [2]int{1, 2}, 2; expectedVal != result || expectedNums != [2]int(nums) {
		t.Errorf("Result must be [%v] with nums [%v] for starts val [%v], actual is [%v] for nums [%v]", expectedVal, expectedNums, val, result, nums)
	}
}

func TestRemovesManyValForStartIntoNums(t *testing.T) {
	nums, val := []int{0, 0, -2, -1}, 0
	result := removeElement(nums, val)
	if expectedNums, expectedVal := [2]int{-2, -1}, 2; expectedVal != result || expectedNums != [2]int(nums) {
		t.Errorf("Result must be [%v] with nums [%v] for starts val [%v], actual is [%v] for nums [%v]", expectedVal, expectedNums, val, result, nums)
	}
}

func TestRemovesFirstAndLastNumFromNums(t *testing.T) {
	nums, val := []int{3, 2, 2, 3}, 3
	result := removeElement(nums, val)
	if expectedNums, expectedVal := [2]int{2, 2}, 2; expectedVal != result || expectedNums != [2]int(nums) {
		t.Errorf("Result must be [%v] with nums [%v] for val [%v], actual is [%v] for nums [%v]", expectedVal, expectedNums, val, result, nums)
	}
}
