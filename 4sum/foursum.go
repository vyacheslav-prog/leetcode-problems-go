package foursum

import "sort"

func findUniqueTargetPairsForSortedNums(nums []int, startIndex, target int) [][]int {
	var result [][]int
	leftIndex, rightIndex := startIndex, len(nums)-1
	for leftIndex < rightIndex {
		if pairSum := nums[leftIndex] + nums[rightIndex]; target < pairSum {
			rightIndex -= 1
		} else if pairSum < target {
			leftIndex += 1
		} else {
			result = append(result, []int{nums[leftIndex], nums[rightIndex]})
			leftIndex += 1
			rightIndex -= 1
		}
	}
	return result
}

func findUniqueTargetSumGroupsForSortedNumsBySize(nums []int, size, startIndex, target int) [][]int {
	if (len(nums) < size) || (target < (nums[0] * size)) || ((nums[len(nums)-1] * size) < target) {
		return nil
	}
	if 2 == target {
		return findUniqueTargetPairsForSortedNums(nums, startIndex, target)
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
