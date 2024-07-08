package foursum

import "testing"

func TestFindsZeroQuadrupletsForZero(t *testing.T) {
	result := fourSum(nil, 0)
	if 0 != len(result) {
		t.Errorf("Result must be empty for empty nums, actual is [%v]", result)
	}
}

func TestFindsSingleQuadrupletWhenSumIsEqualToTarget(t *testing.T) {
	nums, target := []int{1, 1, 1, 1}, 4
	result := fourSum(nums, target)
	if expected := [4]int(nums); 1 != len(result) || expected != [4]int(result[0]) {
		t.Errorf("Result must be [%v] for single quadruplet into nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestFindsZeroQuadrupletsWhenTargetIsNotEqualANumsSum(t *testing.T) {
	nums, target := []int{1, 1, 1, 1}, 5
	result := fourSum(nums, target)
	if 0 != len(result) {
		t.Errorf("Result must be empty for nums [%v] and target [%v], actual is [%v]", nums, target, result)
	}
}

func TestFindsQuadrupletForStartingTargetNumsAndAddsIntoEnd(t *testing.T) {
	nums, target := []int{1, 1, 1, 1, 2}, 4
	result := fourSum(nums, target)
	if expected := [4]int(nums[:4]); 1 != len(result) || expected != [4]int(result[0]) {
		t.Errorf("Result must be [%v] for starting quadruplet into nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestFindsQuadrupletForNumsWithUnsuitableNumIntoMiddle(t *testing.T) {
	nums, target := []int{1, 0, 1, 1, 1}, 4
	result := fourSum(nums, target)
	if expected := [4]int{1, 1, 1, 1}; 1 != len(result) || expected != [4]int(result[0]) {
		t.Errorf("Result must be [%v] for nums [%v] with unsuitable nums and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestFindsQuadrupletWithSkippedNumIntoMiddle(t *testing.T) {
	nums, target := []int{-1, -1, 0, 1, 1}, 0
	result := fourSum(nums, target)
	if expected := [4]int{1, 1, -1, -1}; 1 != len(result) || expected != [4]int(result[0]) {
		t.Errorf("Result must be [%v] for nums [%v] with an excess and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestFindsManyQuadrupletsForZeroTargetAndVariousNums(t *testing.T) {
	nums, target := []int{-2, -1, 0, 0, 1, 2}, 0
	result := fourSum(nums, target)
	if 3 != len(result) {
		t.Errorf("Result must have [3] quadruplets for nums [%v] and target [%v], actual is [%v]", nums, target, result)
	}
}
