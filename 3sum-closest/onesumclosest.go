package threesumclosest

import "sort"

func oneSumClosest(nums []int, target int) int {
	if 0 == len(nums) {
		return 0
	}
	sort.Ints(nums)
	result := nums[0]
	for index, reverse := 0, -1*target; index != len(nums); index += 1 {
		if 0 == index || (reverse+nums[index] <= reverse+result) {
			result = nums[index]
		}
	}
	return result
}
