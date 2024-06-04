package threesum

func threeSum(nums []int) [][]int {
	var result [][]int
	if 2 < len(nums) {
		var sum int
		for _, number := range nums[:3] {
			sum += number
		}
		if 0 == sum {
			result = append(result, nums)
		}
	}
	return result
}
