package searchinrotatedsortedarray

import "log"

const noIndex = -1

func search(nums []int, target int) int {
	if len(nums) < 1 || (1 == len(nums) && target != nums[0]) {
		return noIndex
	}
	middleIndex := len(nums) / 2
	if target == nums[middleIndex] {
		return middleIndex
	}
	if isRotated, middleIsLess := nums[len(nums)-1] < nums[0], nums[middleIndex] < target; (isRotated && middleIsLess != true) || (isRotated != true && middleIsLess) {
		leftIndex := middleIndex
		if 0 == (leftIndex+len(nums))%2 {
			leftIndex += 1
		}
		log.Printf("Right [%v] for nums [%v]", nums[leftIndex:len(nums)], nums)
		nestedResult := search(nums[leftIndex:len(nums)], target)
		if noIndex == nestedResult {
			return noIndex
		}
		return leftIndex + nestedResult
	}
	log.Printf("Left [%v] for nums [%v]", nums[:middleIndex], nums)
	return search(nums[:middleIndex], target)
}
