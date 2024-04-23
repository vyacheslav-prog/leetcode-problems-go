package containerwithmostwater

import "testing"

func TestEmptyWaterForEmptyContainer(t *testing.T) {
	result := maxArea([]int{})
	if result != 0 {
		t.Errorf("Result must be zero for empty container, actual is [%v]", result)
	}
}
