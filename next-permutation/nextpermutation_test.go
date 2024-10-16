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
	nums, result := []int{1, 2, 3, 2, 2}, []int{1, 2, 3, 2, 2}
	nextPermutation(result)
	if expected := [5]int{1, 3, 2, 2, 2}; 5 != len(result) || expected != [5]int(result) {
		t.Errorf("Result must be [%v] for nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestPermutatesWithDescendingRightTwoNums(t *testing.T) {
	nums, result := []int{10, 13, 7}, []int{10, 13, 7}
	nextPermutation(result)
	if expected := [3]int{13, 7, 10}; 3 != len(result) || expected != [3]int(result) {
		t.Errorf("Result must be [%v] for nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestPermutatesWithDescendingRightNums(t *testing.T) {
	nums, result := []int{3, 9, 6, 3}, []int{3, 9, 6, 3}
	nextPermutation(result)
	if expected := [4]int{6, 3, 3, 9}; 4 != len(result) || expected != [4]int(result) {
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

func TestPermutatesNoForSingletonNums(t *testing.T) {
	nums := []int{1}
	nextPermutation(nums)
	if 1 != len(nums) {
		t.Errorf("Result must have single number for singleton nums [%v]", nums)
	}
}

func TestPermutatesLeftDescendingForZeroIndexLeader(t *testing.T) {
	nums, result := []int{3, 1, 2}, []int{3, 1, 2}
	nextPermutation(result)
	if expected := [3]int{3, 2, 1}; 3 != len(result) || expected != [3]int(result) {
		t.Errorf("Result must be [%v] for nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestPermutatesWithNotFirstDescendingAndLongOrderingTail(t *testing.T) {
	nums, result := []int{5, 4, 7, 5, 3, 2}, []int{5, 4, 7, 5, 3, 2}
	nextPermutation(result)
	if expected := [6]int{5, 5, 2, 3, 4, 7}; 6 != len(result) || expected != [6]int(result) {
		t.Errorf("Result must be [%v] for nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestPermutatesLongDistanceForRevertNumberIntoDescendingSequence(t *testing.T) {
	nums, result := []int{2, 2, 7, 5, 4, 3, 2, 2, 1}, []int{2, 2, 7, 5, 4, 3, 2, 2, 1}
	nextPermutation(result)
	if expected := [9]int{2, 3, 1, 2, 2, 2, 4, 5, 7}; 9 != len(result) || expected != [9]int(result) {
		t.Errorf("Result must be [%v] for nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestPermutatesFourDecreasingNums(t *testing.T) {
	nums, result := []int{4, 3, 2, 1}, []int{4, 3, 2, 1}
	nextPermutation(result)
	if expected := [4]int{1, 2, 3, 4}; 4 != len(result) || expected != [4]int(result) {
		t.Errorf("Result must be [%v] for decreasing nums [%v], actual is [%v]", expected, nums, result)
	}
}
