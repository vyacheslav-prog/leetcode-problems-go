package nextpermutation

import "testing"

func TestReplacesZeroNumsForNilNums(t *testing.T) {
	var nums []int
	nextPermutation(nums)
	if nil != nums {
		t.Errorf("Result must be nil for empty nums, actual is [%v]", nums)
	}
}

func TestPermutatesForTwoNums(t *testing.T) {
	nums := []int{1, 2}
	result := nums
	nextPermutation(nums)
	if expected := [2]int{2, 1}; 2 != len(result) || expected != [2]int(result) {
		t.Errorf("Result must be [%v] for two nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestPermutatesForRightNumsForThreeNums(t *testing.T) {
	nums := []int{5, 10, 15}
	result := nums
	nextPermutation(nums)
	if expected := [3]int{5, 15, 10}; 3 != len(result) || expected != [3]int(result) {
		t.Errorf("Result must be [%v] for nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestPermutatesForPairNumsIntoMiddle(t *testing.T) {
	nums, result := []int{1, 2, 3, 2, 1}, []int{1, 2, 3, 2, 1}
	nextPermutation(result)
	if expected := [5]int{1, 3, 2, 2, 1}; 5 != len(result) || expected != [5]int(result) {
		t.Errorf("Result must be [%v] for nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestPermutatesForEndOfIncreasingNums(t *testing.T) {
	nums, result := []int{3, 2, 1}, []int{3, 2, 1}
	nextPermutation(result)
	if expected := [3]int{1, 2, 3}; 3 != len(result) || expected != [3]int(result) {
		t.Errorf("Result must be [%v] for nums [%v], actual is [%v]", expected, nums, result)
	}
}
