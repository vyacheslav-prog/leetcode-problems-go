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

func TestFindsFirstTripletForFourNumsWhenTotalSumIsNotZero(t *testing.T) {
	nums := []int{0, 0, 0, 1}
	result := threeSum(nums)
	if expected := [3]int{0, 0, 0}; 1 != len(result) || expected != [3]int(result[0]) {
		t.Errorf("Result must have [%v] for first triplet for nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestFindsLastTripletForFourNumsWhenTotalSumIsNotZero(t *testing.T) {
	nums := []int{1, 0, 0, 0}
	result := threeSum(nums)
	if expected := [3]int{0, 0, 0}; 1 != len(result) || expected != [3]int(result[0]) {
		t.Errorf("Result must have [%v] for last triplet for nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestFindsMiddleTripletForFiveNumsWhenTotalSumIsNotZero(t *testing.T) {
	nums := []int{2, -1, 0, 1, 2}
	result := threeSum(nums)
	if expected := [3]int{-1, 0, 1}; 1 != len(result) || expected != [3]int(result[0]) {
		t.Errorf("Result must have [%v] for middle triplet for nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestFindsOutOfOrderTriplet(t *testing.T) {
	nums := []int{-1, 0, 0, 1}
	result := threeSum(nums)
	if expected := [3]int{-1, 0, 1}; 1 != len(result) || expected != [3]int(result[0]) {
		t.Errorf("Result must have [%v] for out of order triplet for nums [%v], actual is [%v]", expected, nums, result)
	}
}
