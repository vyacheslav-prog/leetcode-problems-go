package searchinrotatedsortedarray

import "fmt"

const noIndex = -1

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return noIndex
	}
	middleIndex := len(nums) / 2
	if isOdd := 0 != len(nums)%2; isOdd && 0 != middleIndex {
		middleIndex += 1
	}
	if target == nums[middleIndex] {
		return middleIndex
	}
	if isRotated := nums[len(nums)-1] < nums[0]; isRotated || nums[middleIndex] < target {
		fmt.Printf("right [%v] for nums [%v] and target [%v]\n", nums[middleIndex:len(nums)-1], nums, target)
		nestedResult := search(nums[middleIndex:len(nums)-1], target)
		if noIndex == nestedResult {
			return noIndex
		}
		return middleIndex + nestedResult
	}
	fmt.Printf("left [%v] for nums [%v] and target [%v]\n", nums[:middleIndex], nums, target)
	return search(nums[:middleIndex], target)
}
