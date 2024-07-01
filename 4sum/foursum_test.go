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
