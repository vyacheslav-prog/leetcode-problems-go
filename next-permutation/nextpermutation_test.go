package nextpermutation

import "testing"

func TestReplacesZeroNumsForNilNums(t *testing.T) {
	var nums []int
	nextPermutation(nums)
	if nil != nums {
		t.Errorf("Result must be nil for empty nums, actual is [%v]", nums)
	}
}
