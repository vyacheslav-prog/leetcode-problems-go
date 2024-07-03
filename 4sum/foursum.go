package foursum

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	var result [][]int
	var sum int
	for index := 3; index != len(nums); index += 1 {
		sum = nums[index] + nums[index-1] + nums[index-2] + nums[index-3]
		if target == sum {
			result = append(result, nums[index-3:index+1])
		}
	}
	return result
}
