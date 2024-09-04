package removeelement

func removeElement(nums []int, val int) int {
	var valCounter int
	for index := 0; len(nums) != index; index += 1 {
		if val == nums[index] {
			nums = append(nums[:index], nums[index+1:]...)
			index -= 1
			valCounter += 1
		}
	}
	return len(nums)
}
