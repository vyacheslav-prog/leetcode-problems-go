package searchinrotatedsortedarray

const noIndex = -1

func search(nums []int, target int) int {
	if len(nums) < 1 || (1 == len(nums) && target != nums[0]) {
		return noIndex
	}
	middleIndex := len(nums) / 2
	if isOdd := 0 != len(nums)%2; isOdd && 0 != middleIndex {
		middleIndex += 1
	}
	if target == nums[middleIndex] {
		return middleIndex
	}
	if isRotated := nums[len(nums)-1] < nums[0]; (isRotated && target < nums[0]) || (isRotated != true && nums[middleIndex] < target) {
		nestedResult := search(nums[middleIndex:len(nums)], target)
		if noIndex == nestedResult {
			return noIndex
		}
		return middleIndex + nestedResult
	}
	return search(nums[:middleIndex], target)
}
