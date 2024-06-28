package foursum

import "testing"

func TestFindsZeroQuadrupletsForZero(t *testing.T) {
	result := fourSum(nil, 0)
	if 0 != len(result) {
		t.Errorf("Result must be empty for empty nums, actual is [%v]", result)
	}
}
