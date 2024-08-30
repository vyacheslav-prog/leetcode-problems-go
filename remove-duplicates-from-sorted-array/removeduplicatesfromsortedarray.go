package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	var uniqueCounter int
	for counter := 0; len(nums) != counter; counter += 1 {
		if 0 == counter || nums[counter-1] != nums[counter] {
			uniqueCounter += 1
		}
	}
	return uniqueCounter
}
