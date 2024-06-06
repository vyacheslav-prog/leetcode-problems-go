package threesum

func threeSum(nums []int) [][]int {
	var result [][]int
	var triplet []int
	for firstIndex := 0; firstIndex != len(nums) && 0 == len(triplet); firstIndex += 1 {
		for secondIndex := firstIndex + 1; secondIndex != len(nums); secondIndex += 1 {
			for balance, thirdIndex := nums[firstIndex]+nums[secondIndex], secondIndex+1; thirdIndex != len(nums); thirdIndex += 1 {
				if 0 == nums[thirdIndex]+balance {
					triplet = []int{nums[firstIndex], nums[secondIndex], nums[thirdIndex]}
					break
				}
			}
		}
	}
	if 3 == len(triplet) {
		result = append(result, triplet)
	}
	return result
}
