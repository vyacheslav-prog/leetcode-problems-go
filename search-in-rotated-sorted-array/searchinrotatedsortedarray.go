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
	if target == nums[middleIndex] {
		return middleIndex
	}
	log.Printf("For nums [%v] and target [%v], middle [%v]", nums, target, nums[middleIndex])
	if isRotated, middleIsLess := nums[len(nums)-1] < nums[0], nums[middleIndex] < target; (isRotated != true && middleIsLess) || (isRotated && nums[len(nums)-1] < nums[middleIndex] && target != nums[0]) {
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
