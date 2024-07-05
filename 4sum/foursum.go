package foursum

import "sort"

func findUniqueTargetPairsForSortedNums(nums []int, target int) [][]int {
	var result [][]int
	firstIndex, lastIndex := 0, len(nums)-1
	for firstIndex < lastIndex {
		if target == (nums[firstIndex] + nums[lastIndex]) {
			result = append(result, []int{nums[firstIndex], nums[lastIndex]})
			firstIndex += 1
			lastIndex -= 1
		} else if target < (nums[firstIndex] + nums[lastIndex]) {
			firstIndex += 1
		} else {
			lastIndex -= 1
		}
	}
	return result
}

func findUniqueTargetSumGroupsForSortedNumsBySize(nums []int, size, startIndex, target int) [][]int {
	if len(nums) < size {
		return nil
	}
	if (target / size) < nums[0] {
		return nil
	}
	if nums[len(nums)-1] < (target / size) {
		return nil
	}
	if 2 == target {
		return findUniqueTargetPairsForSortedNums(nums, target)
	}
	var result [][]int
	for index := startIndex; index != len(nums); index += 1 {
		if index == startIndex || nums[index] != nums[index-1] {
			for _, subGroup := range findUniqueTargetSumGroupsForSortedNumsBySize(nums, size-1, index+1, target-nums[index]) {
				result = append(result, append(subGroup, nums[index]))
			}
		}
	}
	return result
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return findUniqueTargetSumGroupsForSortedNumsBySize(nums, 4, 0, target)
}
