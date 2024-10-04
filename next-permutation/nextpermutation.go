package nextpermutation

func nextPermutation(nums []int) {
	if 2 == len(nums) {
		nums[0], nums[1] = nums[1], nums[0]
	}
}
