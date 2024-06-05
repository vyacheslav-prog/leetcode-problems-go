package threesum

func threeSum(nums []int) [][]int {
	var result [][]int
	var lastIndex, sum int
	for index := 0; index != len(nums) && (lastIndex < 3 || 0 != sum); index += 1 {
		if 2 < lastIndex {
			sum -= nums[lastIndex-3]
		}
		sum += nums[index]
		lastIndex += 1
	}
	if 0 == sum && 2 < lastIndex {
		result = append(result, nums[lastIndex-3:lastIndex])
	}
	return result
}
