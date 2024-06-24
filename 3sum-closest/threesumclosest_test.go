package threesumclosest

import "testing"

func TestFindsZeroForNilNums(t *testing.T) {
	result := threeSumClosest(nil, 0)
	if 0 != result {
		t.Errorf("Result must be zero for nil nums, actual is [%v]", result)
	}
}

func TestFindsSumForThreeNums(t *testing.T) {
	nums := []int{1, 1, 1}
	result := threeSumClosest(nums, 0)
	if expected := 3; expected != result {
		t.Errorf("Result must be [%v] for three nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestFindsSumThreeForFourNums(t *testing.T) {
	nums, target := []int{0, 1, 1, 1}, 0
	result := threeSumClosest(nums, target)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for four nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestFindsMaxSumThreeForFourNums(t *testing.T) {
	nums, target := []int{0, 1, 1, 1}, 3
	result := threeSumClosest(nums, target)
	if expected := 3; expected != result {
		t.Errorf("Result must be [%v] for nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestFindsClosestSumOfThreeWhenNumsIsUnordered(t *testing.T) {
	nums, target := []int{1, 1, 1, 0}, 3
	result := threeSumClosest(nums, target)
	if expected := 3; expected != result {
		t.Errorf("Result must be [%v] for unordered nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestFindsClosestRightOverflowedSumOfThree(t *testing.T) {
	nums, target := []int{-1, 2, 1, -4}, 1
	result := threeSumClosest(nums, target)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for overflowed sum of nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestFindsClosestSumForFirstAndMinimalTriplet(t *testing.T) {
	nums, target := []int{2, 3, 8, 9, 10}, 16
	result := threeSumClosest(nums, target)
	if expected := 15; expected != result {
		t.Errorf("Result must be minimal [%v] for nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}
