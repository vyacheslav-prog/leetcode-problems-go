package threesum

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	var prevTriplet, triplet []int
	sort.Ints(nums)
	for firstIndex := 0; firstIndex != len(nums); firstIndex += 1 {
		for secondIndex := firstIndex + 1; secondIndex != len(nums); secondIndex += 1 {
			balance := nums[firstIndex] + nums[secondIndex]
			for thirdIndex := secondIndex + 1; thirdIndex != len(nums); thirdIndex += 1 {
				if 0 == nums[thirdIndex]+balance {
					triplet = []int{nums[firstIndex], nums[secondIndex], nums[thirdIndex]}
					break
				}
			}
			if 3 == len(triplet) && (nil == prevTriplet || [3]int(prevTriplet) != [3]int(triplet)) {
				result = append(result, triplet)
				prevTriplet = triplet
				triplet = nil
			}
		}
	}
	return result
}
