package threesumclosest

import "sort"

func threeSumClosest(nums []int, target int) int {
	var distance, result, resultDistance, sum int
	sort.Ints(nums)
	for firstIndex := 0; firstIndex < len(nums)-2; firstIndex += 1 {
		for secondIndex, thirdIndex := firstIndex+1, len(nums)-1; secondIndex != thirdIndex; {
			sum = nums[firstIndex] + nums[secondIndex] + nums[thirdIndex]
			distance = target - sum
			if distance < 0 {
				distance *= -1
			}
			if result == 0 || distance < resultDistance {
				result = sum
				resultDistance = distance
			}
			if sum < target {
				secondIndex += 1
			} else if target < sum {
				thirdIndex -= 1
			} else {
				secondIndex = thirdIndex
			}
		}
	}
	return result
}
