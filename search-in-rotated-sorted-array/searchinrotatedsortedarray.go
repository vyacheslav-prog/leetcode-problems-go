package searchinrotatedsortedarray

const noIndex = -1

func search(nums []int, target int) int {
	if len(nums) < 1 || (1 == len(nums) && target != nums[0]) {
		return noIndex
	}
	middleIndex := len(nums) / 2
	if target == nums[middleIndex] {
		return middleIndex
	}
	if isRotated := nums[len(nums)-1] < nums[0]; (isRotated && target < nums[0]) || (isRotated != true && nums[middleIndex] < target) {
		leftIndex := middleIndex
		if 0 == (leftIndex+len(nums))%2 {
			leftIndex += 1
		}
		nestedResult := search(nums[leftIndex:len(nums)], target)
		if noIndex == nestedResult {
			return noIndex
		}
		return leftIndex + nestedResult
	}
	return search(nums[:middleIndex], target)
}
