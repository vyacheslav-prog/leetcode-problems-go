package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	var nextUniqueIndex, uniqueCounter int
	for counter := 0; len(nums) != counter; counter += 1 {
		nextUniqueIndex = counter + 1
		for len(nums) != nextUniqueIndex && nums[counter] == nums[nextUniqueIndex] {
			nextUniqueIndex += 1
		}
		for swapIndex := nextUniqueIndex - 1; len(nums) != nextUniqueIndex && counter != swapIndex; swapIndex -= 1 {
			nums[swapIndex] = nums[nextUniqueIndex]
		}
		if 0 == counter || nums[counter-1] != nums[counter] {
			uniqueCounter += 1
		}
	}
	return uniqueCounter
}
