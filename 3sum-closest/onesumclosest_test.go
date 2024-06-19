package threesumclosest

import "testing"

func TestOneFindsZeroForNilNums(t *testing.T) {
	result := oneSumClosest(nil, 0)
	if 0 != result {
		t.Errorf("Result must be zero for nil nums, actual is [%v]", result)
	}
}

func TestOneFindsFirstNumForSingletonNums(t *testing.T) {
	nums := []int{-1}
	result := oneSumClosest(nums, 0)
	if expected := -1; expected != result {
		t.Errorf("Result must be [%v] for singleton nums [%v], actual is [%v]", expected, nums, result)
	}
}

func TestOneFindsNumForEqualedTarget(t *testing.T) {
	nums, target := []int{-1, 0, 1}, 0
	result := oneSumClosest(nums, target)
	if expected := 0; expected != result {
		t.Errorf("Result must be [%v] for found target [%v] into nums [%v], actual is [%v]", expected, target, nums, result)
	}
}
