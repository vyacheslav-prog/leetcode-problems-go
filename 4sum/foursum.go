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
			if (startIndex == leftIndex || nums[leftIndex] != nums[leftIndex-1]) && (rightIndex == (len(nums)-1) || nums[rightIndex] != nums[rightIndex+1]) {
				result = append(result, []int{nums[leftIndex], nums[rightIndex]})
			}
			leftIndex += 1
			rightIndex -= 1
		}
	}
	return result
}

func findUniqueTargetSumGroupsForSortedNumsBySize(nums []int, size, startIndex, target int) [][]int {
	var result [][]int
	if (len(nums) < size) || (target < (nums[startIndex] * size)) || ((nums[len(nums)-1] * size) < target) {
		return result
	}
	if 2 == size {
		return findUniqueTargetPairsForSortedNums(nums, startIndex, target)
	}
	for index := startIndex; index != len(nums)-1; index += 1 {
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
