package searchinrotatedsortedarray

const noIndex = -1

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return noIndex
	}
	middleIndex := len(nums) / 2
	if target == nums[middleIndex] {
		return middleIndex
	}
	if nums[middleIndex] < target && nums[0] < nums[len(nums)-1] {
		return search(nums[middleIndex:len(nums)-1], target)
	}
	return search(nums[0:middleIndex], target)
}
