package nextpermutation

func nextPermutation(nums []int) {
	descendingNumsIndex, leaderSwapIndex := -1, -1
	for descendingNumsValue, index := 0, 0; len(nums) != index; index += 1 {
		if -1 == descendingNumsIndex || descendingNumsValue < nums[index] || nums[index-1] < nums[index] {
			descendingNumsIndex = index
			descendingNumsValue = nums[index]
		}
	}
	if 0 == descendingNumsIndex && 1 != len(nums) {
		descendingNumsIndex = 1
		leaderSwapIndex = len(nums) - 1
	} else if -1 != descendingNumsIndex {
		leaderSwapIndex = descendingNumsIndex - 1
		for index, swapIndex := descendingNumsIndex+1, leaderSwapIndex; len(nums) != index && nums[swapIndex] < nums[index]; index += 1 {
			leaderSwapIndex = index
		}
	}
	if -1 != leaderSwapIndex {
		if leaderSwapIndex == descendingNumsIndex-1 {
			leaderSwapIndex = descendingNumsIndex
		}
		nums[descendingNumsIndex-1], nums[leaderSwapIndex] = nums[leaderSwapIndex], nums[descendingNumsIndex-1]
	}
	for endIndex := len(nums) - 1; descendingNumsIndex != endIndex; endIndex -= 1 {
		for index := descendingNumsIndex; index < endIndex; index += 1 {
			if nums[index+1] < nums[index] {
				nums[index+1], nums[index] = nums[index], nums[index+1]
			}
		}
	}
}
