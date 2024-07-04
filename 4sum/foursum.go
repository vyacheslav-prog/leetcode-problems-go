package foursum

import "sort"

const size int = 4

func fourSum(nums []int, target int) [][]int {
	if len(nums) < size {
		return nil
	}
	var quadruplet []int
	var quadrupletSum int
	var result [][]int
	sort.Ints(nums)
	for index := size - 1; index != len(nums); index += 1 {
		quadruplet = nums[index-size+1 : index+1]
		quadrupletSum = 0
		for _, value := range quadruplet {
			quadrupletSum += value
		}
		if target == quadrupletSum {
			result = append(result, quadruplet)
		}
	}
	return result
}
