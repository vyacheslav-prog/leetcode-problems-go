package searchinrotatedsortedarray

import "testing"

func TestSearchesNoIndexForEmptyArray(t *testing.T) {
	result := search(nil, 0)
	if noIndex != result {
		t.Errorf("Result must be -1 for empty array, actual is [%v]", result)
	}
}

func TestSearchesIndexForSingletonArray(t *testing.T) {
	nums, target := []int{1}, 1
	result := search(nums, target)
	if expected := 0; expected != result {
		t.Errorf("Result must be [%v] for singleton array [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestSearchesNoIndexWhenTargetIsNotPresented(t *testing.T) {
	nums, target := []int{1}, 2
	result := search(nums, target)
	if noIndex != result {
		t.Errorf("Result must be -1 for not empty nums [%v] and not presented target [%v], actual is [%v]", nums, target, result)
	}
}

func TestSearchesIndexForLastTargetNum(t *testing.T) {
	nums, target := []int{2, 3}, 3
	result := search(nums, target)
	if expected := 1; expected != result {
		t.Errorf("Result must be [%v] for nums [%v] and target [%v] into last position, actual is [%v]", expected, nums, target, result)
	}
}

func TestSearchesIndexForLastTargetElementWhenTwoElementsNumsIsRotated(t *testing.T) {
	nums, target := []int{2, -2}, 2
	result := search(nums, target)
	if expected := 0; expected != result {
		t.Errorf("Result must be [%v] for rotated nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestSearchesIndexForLastTargetElementWhenThreeElementsNumsIsRotated(t *testing.T) {
	nums, target := []int{1, 2, 0}, 0
	result := search(nums, target)
	if expected := 2; expected != result {
		t.Errorf("Result must be [%v] for rotated nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestSearchesIndexForTwoStepSearchingForNotRotatedNums(t *testing.T) {
	nums, target := []int{-5, -3, -1, 1, 3, 5}, 3
	result := search(nums, target)
	if expected := 4; expected != result {
		t.Errorf("Result must be [%v] for nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestSearchesIndexForTwoStepSearchingForRotatedNums(t *testing.T) {
	nums, target := []int{-3, -1, 1, 3, 5, -5}, 3
	result := search(nums, target)
	if expected := 3; expected != result {
		t.Errorf("Result must be [%v] for nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestSearchesForTwoNumsWithoutTarget(t *testing.T) {
	nums, target := []int{1, 3}, 4
	result := search(nums, target)
	if noIndex != result {
		t.Errorf("Result must be -1 for nums [%v] and not presented target [%v], actual is [%v]", nums, target, result)
	}
}

func TestSearchesIndexForThreeStepSearchingForRotatedNums(t *testing.T) {
	nums, target := []int{4, 5, 6, 7, 0, 1, 2}, 1
	result := search(nums, target)
	if expected := 5; expected != result {
		t.Errorf("Result must be [%v] for nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestSearchesIndexForMiddleNumForRotatedNums(t *testing.T) {
	nums, target := []int{5, 1, 3}, 1
	result := search(nums, target)
	if expected := 1; expected != result {
		t.Errorf("Result must be [%v] for nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestSearchedIndexForRotatedNumsWhenTargetIsNextFromMiddle(t *testing.T) {
	nums, target := []int{4, 5, 6, 7, 0, 1, 2}, 0
	result := search(nums, target)
	if expected := 4; expected != result {
		t.Errorf("Result must be [%v] for nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}

func TestSearchedIndexForRotatedNumsWhenTargetIsPrevFromMiddle(t *testing.T) {
	nums, target := []int{7, 8, 1, 2, 3, 4, 5, 6}, 2
	result := search(nums, target)
	if expected := 3; expected != result {
		t.Errorf("Result must be [%v] for nums [%v] and target [%v], actual is [%v]", expected, nums, target, result)
	}
}
