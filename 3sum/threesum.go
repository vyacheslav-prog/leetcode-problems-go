package threesum

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for firstIndex := 0; firstIndex != len(nums); firstIndex += 1 {
		firstNumber := nums[firstIndex]
		if firstIndex != 0 && firstNumber == nums[firstIndex-1] {
			continue
		}
		balance := -1 * firstNumber
		leftIndex, rightIndex := firstIndex+1, len(nums)-1
		for leftIndex < rightIndex {
			leftNumber, rightNumber := nums[leftIndex], nums[rightIndex]
			if twoSum := leftNumber + rightNumber; balance == twoSum {
				triplet := []int{firstNumber, leftNumber, rightNumber}
				if 0 == len(result) || [3]int(result[len(result)-1]) != [3]int(triplet) {
					result = append(result, triplet)
				}
				leftIndex += 1
				rightIndex -= 1
			} else if twoSum < balance {
				leftIndex += 1
			} else {
				rightIndex -= 1
			}
		}
	}
	return result
}
