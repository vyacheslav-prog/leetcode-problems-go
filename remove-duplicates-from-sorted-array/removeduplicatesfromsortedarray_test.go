package removeduplicatesfromsortedarray

import "testing"

func TestRemovesZeroNumsWhenNumsIsEmpty(t *testing.T) {
	var nums []int
	result := removeDuplicates(nums)
	if 0 != result || nil != nums {
		t.Errorf("Result must be zero k-nums and nil for empty nums array, actual is [%v]", result)
	}
}

func TestRemovesZeroNumsWhenNumsIsUnique(t *testing.T) {
	nums := []int{0}
	result := removeDuplicates(nums)
	if expected := len(nums); expected != result || expected != len(nums) {
		t.Errorf("Result must be [%v] and nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestRemovesAllDuplicatesForRepeatedNums(t *testing.T) {
	nums := []int{1, 1}
	result := removeDuplicates(nums)
	if expected := 1; expected != result || 1 != nums[0] {
		t.Errorf("Result must be [%v], actual is [%v] and [%v]", expected, result, nums)
	}
}

func TestRemovesDuplicatesForNumsWithOnePairedNums(t *testing.T) {
	nums := []int{1, 1, 2}
	result := removeDuplicates(nums)
	if expected, expectedNums := 2, [2]int{1, 2}; expected != result || expectedNums != [2]int(nums) {
		t.Errorf("Result must be [%v], actual is [%v] with nums [%v]", expected, result, nums)
	}
}

func TestRemovesDuplicatesForNumsWithOnePairedNumsInEnd(t *testing.T) {
	nums := []int{3, 4, 4}
	result := removeDuplicates(nums)
	if expected, expectedNums := 2, [2]int{3, 4}; expected != result || expectedNums != [2]int(nums) {
		t.Errorf("Result must be [%v], actual is [%v] with nums [%v]", expected, result, nums)
	}
}

func TestRemovesAllDuplicatesWithUniqueEndsNum(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(nums)
	if expected, expectedNums := 5, [5]int{0, 1, 2, 3, 4}; expected != result || expectedNums != [5]int(nums) {
		t.Errorf("Result must be [%v], actual is [%v] with nums [%v]", expected, result, nums)
	}
}
