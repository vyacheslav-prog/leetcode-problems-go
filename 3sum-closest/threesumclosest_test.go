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
	nums := []int{0, 1, 1, 1}
	result := threeSumClosest(nums, 0)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for four nums [%v], actual is [%v]", expected, nums, result)
	}
}
