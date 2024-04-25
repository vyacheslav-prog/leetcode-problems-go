package containerwithmostwater

import "testing"

func TestEmptyWaterForEmptyContainer(t *testing.T) {
	result := maxArea([]int{})
	if result != 0 {
		t.Errorf("Result must be zero for empty container, actual is [%v]", result)
	}
}

func TestStoresOneAmountForTwoSingleHeightLines(t *testing.T) {
	height := []int{1, 1}
	result := maxArea(height)
	if expected := 1; result != expected {
		t.Errorf("Result must be [%v] for two single-height lines [%v], actual is [%v]", expected, height, result)
	}
}

func TestStoresZeroAmountForOneLine(t *testing.T) {
	height := []int{2}
	result := maxArea(height)
	if result != 0 {
		t.Errorf("Result must be zero for one line [%v], actual is [%v]", height, result)
	}
}

func TestStoresOneAmountForTwoDescendingLinesWithMinimumOfOne(t *testing.T) {
	height := []int{2, 1}
	result := maxArea(height)
	if expected := 1; result != expected {
		t.Errorf("Result must be [%v] for a descending order lines [%v], actual is [%v]", expected, height, result)
	}
}

func TestStoresOneAmountForTwoAscendingLinesWithMinimumOfOne(t *testing.T) {
	height := []int{1, 2}
	result := maxArea(height)
	if expected := 1; result != expected {
		t.Errorf("Result must be [%v] for a ascending order lines [%v], actual is [%v]", expected, height, result)
	}
}

func TestStoresTwoAmountForThreeSingleHeightLines(t *testing.T) {
	height := []int{1, 1, 1}
	result := maxArea(height)
	if expected := 2; result != expected {
		t.Errorf("Result must be [%v] for three single-height lines [%v], actual is [%v]", expected, height, result)
	}
}

func TestStoresWaterWhenFirstLineIsNotBorderedContainer(t *testing.T) {
	height := []int{1, 3, 3}
	result := maxArea(height)
	if expected := 3; result != expected {
		t.Errorf("Result must be [%v] for lines with non-bordered first line [%v], actual is [%v]", expected, height, result)
	}
}

func TestStoresWaterForMiddleHighestLineAndEqualsFirstAndLastLines(t *testing.T) {
	height := []int{1, 2, 1}
	result := maxArea(height)
	if expected := 2; result != expected {
		t.Errorf("Result must be [%v] for lines with peak into middle [%v], actual is [%v]", expected, height, result)
	}
}
