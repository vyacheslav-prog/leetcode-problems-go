package threesumclosest

import "sort"

func oneSumClosest(nums []int, target int) int {
	if 0 == len(nums) {
		return 0
	}
	sort.Ints(nums)
	result := nums[0]
	for index, resultDistance := 0, 0; index != len(nums); index += 1 {
		distance := target - nums[index]
		if distance < 0 {
			distance *= -1
		}
		if 0 == index || distance < resultDistance {
			result = nums[index]
			resultDistance = distance
		}
	}
	return result
}
