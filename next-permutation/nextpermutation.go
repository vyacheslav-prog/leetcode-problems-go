package nextpermutation

func nextPermutation(nums []int) {
	if 1 < len(nums) {
		nums[len(nums)-2], nums[len(nums)-1] = nums[len(nums)-1], nums[len(nums)-2]
	}
}
