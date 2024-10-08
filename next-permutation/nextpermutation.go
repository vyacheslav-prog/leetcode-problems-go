package nextpermutation

func nextPermutation(nums []int) {
	var descendingNumsValue int
	descendingNumsIndex := -1
	for index, value := range nums {
		if -1 == descendingNumsIndex || descendingNumsValue < value {
			descendingNumsIndex = index
			descendingNumsValue = value
		}
	}
	if 0 < descendingNumsIndex {
		nums[descendingNumsIndex-1], nums[descendingNumsIndex] = descendingNumsValue, nums[descendingNumsIndex-1]
	}
}
