package searchinrotatedsortedarray

const noIndex = -1

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return noIndex
	}
	middleIndex := len(nums) / 2
	if target < nums[middleIndex] {
		return search(nums[0:middleIndex], target)
	} else if target == nums[middleIndex] {
		return middleIndex
	}
	return search(nums[middleIndex:len(nums)-1], target)
}
