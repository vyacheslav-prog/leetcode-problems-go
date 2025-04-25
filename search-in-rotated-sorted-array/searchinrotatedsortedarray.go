package searchinrotatedsortedarray

import "log"

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
	case last:
		return lastIndex
	}
	log.Printf("For nums [%v] and target [%v], current [%v] first [%v] last [%v]", nums, target, current, first, last)
	if (first < last && current < target) || (last < current && (target < first || current < target)) {
		log.Printf("Right [%v]", nums[currentIndex:len(nums)])
		nextResult := search(nums[currentIndex:len(nums)], target)
		if noIndex == nextResult {
			return noIndex
		}
		return currentIndex + nextResult
	}
	log.Printf("Left [%v]", nums[:currentIndex])
	return search(nums[:currentIndex], target)
}
