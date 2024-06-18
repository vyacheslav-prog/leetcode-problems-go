package threesumclosest

import "testing"

func TestFindsZeroForNilNums(t *testing.T) {
	result := threeSumClosest(nil, 0)
	if 0 != result {
		t.Errorf("Result must be zero for nil nums, actual is [%v]", result)
	}
}
