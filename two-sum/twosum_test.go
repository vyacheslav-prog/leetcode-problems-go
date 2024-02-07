package twosum

import "testing"

func TestFindsNoIndicesForEmptyNums(t *testing.T) {
	nums := []int{}
	result := TwoSum(nums, 0)
	if len(result) != 0 {
		t.Errorf("Must be zero indices on zero nums, result: %v", result)
	}
}

func TestFindsNoIndicesForSingletonNums(t *testing.T) {
	nums := []int{0}
	target := 0
	result := TwoSum(nums, target)
	if len(result) != 0 {
		t.Errorf("Unexpected indices for one num, result: %v", result)
	}
}

func TestFindsNoIndicesForTargetGreatThanNums(t *testing.T) {
	nums := []int{0}
	target := 1
	result := TwoSum(nums, target)
	if len(result) != 0 {
		t.Errorf("Unexpected indices for zero sum the nums, result: %v", result)
	}
}

func TestFindsNoIndicesWhenNumsIsNotEnough(t *testing.T) {
	nums := []int{1, 1}
	target := 3
	result := TwoSum(nums, target)
	if len(result) != 0 {
		t.Errorf("Result must be empty, result: %v", result)
	}
}

func TestFindsTwoIndicesWhenSumEqualsTarget(t *testing.T) {
	nums := []int{1, 1}
	target := 2
	result := TwoSum(nums, target)
	if result[0] != 0 || result[1] != 1 {
		t.Errorf("Expected two indices when their a sum equals to target, result: %v", result)
	}
}

func TestFindsNoIndicesWhenTheNumsGreatThanTarget(t *testing.T) {
	nums := []int{1}
	target := 0
	result := TwoSum(nums, target)
	if len(result) != 0 {
		t.Errorf("Unexpected indices for the overflow nums, result %v", result)
	}
}

func TestFindsTwoIndicesWhenNumsGreatThanTwo(t *testing.T) {
	nums := []int{1, 1, 1}
	target := 2
	result := TwoSum(nums, target)
	if len(result) != 2 {
		t.Errorf("Unexpected indices when the nums great than two, result: %v", result)
	}
}

func TestFindsIndicesWhenFirstNumNotSuitable(t *testing.T) {
	nums := []int{1, 2, 2}
	target := 4
	result := TwoSum(nums, target)
	if result[0] != 1 || result[1] != 2 {
		t.Errorf("Unexpected indices for the nums when first the num not suitable, result: %v", result)
	}
}

func TestFindsIndicesWhenSecondNumNotSuitable(t *testing.T) {
	nums := []int{2, 1, 2}
	target := 4
	result := TwoSum(nums, target)
	if result[0] != 0 || result[1] != 2 {
		t.Errorf("Unexpected indices for the nums when second the num not suitable, result: %v", result)
	}
}

func TestFindsNoIndicesWhenTargetDoesNotFitsIntoTwoNums(t *testing.T) {
	nums := []int{1, 1, 1}
	target := 3
	result := TwoSum(nums, target)
	if len(result) != 0 {
		t.Errorf("Unexpected indices for nums when target does not fits into two the nums, result %v", result)
	}
}

func TestFindsZerosWhenTargetIsZeroAndNumsWithNoZeros(t *testing.T) {
	nums := []int{0, 4, 3, 0}
	target := 0
	result := TwoSum(nums, target)
	if result[0] != 0 || result[1] != 3 {
		t.Errorf("Unexpected indices for zero target with two a zeros into nums, result %v", result)
	}
}

func TestFindsIndicesWhenNumsAndTargetIsNegative(t *testing.T) {
	nums := []int{-1, -2, -3, -4, -5}
	target := -8
	result := TwoSum(nums, target)
	if result[0] != 2 || result[1] != 4 {
		t.Errorf("Result must be not empty when the nums is negative, result: %v", result)
	}
}
