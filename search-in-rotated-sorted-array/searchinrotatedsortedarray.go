package searchinrotatedsortedarray

const noIndex = -1

func search(nums []int, target int) int {
	currentIndex, lastIndex := len(nums)/2, len(nums)-1
	if lastIndex < 0 {
		return noIndex
	}
	current, first, last := nums[currentIndex], nums[0], nums[lastIndex]
	switch target {
	case current:
		return currentIndex
	case first:
		return 0
	case last:
		return lastIndex
	}
	if (first < last && current < target) || (last < current && (target < first || current < target)) || (current < first && current < target && target < last) {
		nextResult := search(nums[currentIndex:len(nums)], target)
		if noIndex == nextResult {
			return noIndex
		}
		return currentIndex + nextResult
	}
	return search(nums[:currentIndex], target)
}
