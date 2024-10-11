package nextpermutation

func nextPermutation(nums []int) {
	var descendingNumsValue int
	descendingNumsIndex, leaderSwapIndex := -1, -1
	for index, value := range nums {
		if -1 == descendingNumsIndex || descendingNumsValue < value {
			descendingNumsIndex = index
			descendingNumsValue = value
		}
	}
	if 0 < descendingNumsIndex {
		leaderSwapIndex = descendingNumsIndex - 1
	} else if 0 == descendingNumsIndex && 1 != len(nums) {
		descendingNumsIndex = 1
		leaderSwapIndex = len(nums) - 1
	}
	for index := descendingNumsIndex + 1; len(nums) != index; index += 1 {
		if nums[leaderSwapIndex] < nums[index] {
			leaderSwapIndex = index
		}
	}
	if -1 != leaderSwapIndex {
		if leaderSwapIndex == descendingNumsIndex-1 {
			leaderSwapIndex = descendingNumsIndex
		}
		nums[descendingNumsIndex-1], nums[leaderSwapIndex] = nums[leaderSwapIndex], nums[descendingNumsIndex-1]
	}
	for index := descendingNumsIndex; index < len(nums)-1; index += 1 {
		if nums[index+1] < nums[index] {
			nums[index+1], nums[index] = nums[index], nums[index+1]
		}
	}
}
