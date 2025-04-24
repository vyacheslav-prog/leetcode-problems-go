package searchinrotatedsortedarray

import "log"

const noIndex = -1

func search(nums []int, target int) int {
	if len(nums) < 1 || (1 == len(nums) && target != nums[0]) {
		return noIndex
	}
	middleIndex := len(nums) / 2
	if 0 == len(nums)%2 {
		middleIndex -= 1
	}
	current, first, last := nums[middleIndex], nums[0], nums[len(nums)-1]
	if target == current {
		return middleIndex
	}
	if target == last {
		return len(nums) - 1
	}
	log.Printf("For nums [%v] and target [%v], current num is [%v]", nums, target, current)
	if (first < last && current < target) || (last < current && (target < first || current < target)) {
		leftIndex := middleIndex
		if 0 == (leftIndex+len(nums))%2 {
			leftIndex += 1
		}
		log.Printf("Right [%v]", nums[leftIndex:len(nums)])
		nestedResult := search(nums[leftIndex:len(nums)], target)
		if noIndex == nestedResult {
			return noIndex
		}
		return leftIndex + nestedResult
	}
	log.Printf("Left [%v]", nums[:middleIndex])
	return search(nums[:middleIndex], target)
}
