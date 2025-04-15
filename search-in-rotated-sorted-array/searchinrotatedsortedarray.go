package searchinrotatedsortedarray

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
	if isRotated := nums[0] < nums[len(nums)-1]; nums[middleIndex] < target && isRotated {
		nestedResult := search(nums[middleIndex:len(nums)-1], target)
		if noIndex == nestedResult {
			return noIndex
		}
		return middleIndex + nestedResult
	}
	return search(nums[0:middleIndex], target)
}
