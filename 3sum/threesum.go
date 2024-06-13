package threesum

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	var triplet []int
	prevBalances := make(map[int]int)
	sort.Ints(nums)
	for firstIndex := 0; firstIndex != len(nums); firstIndex += 1 {
		for secondIndex := firstIndex + 1; secondIndex != len(nums); secondIndex += 1 {
			balance := nums[firstIndex] + nums[secondIndex]
			triplet = nil
			for thirdIndex := secondIndex + 1; nil == triplet && thirdIndex != len(nums); thirdIndex += 1 {
				if 0 == nums[thirdIndex]+balance {
					triplet = []int{nums[firstIndex], nums[secondIndex], nums[thirdIndex]}
				}
			}
			if firstNumber, alreadyFound := prevBalances[balance]; nil != triplet && (alreadyFound != true || firstNumber != triplet[0]) {
				prevBalances[balance] = triplet[0]
				result = append(result, triplet)
			}
		}
	}
	return result
}
