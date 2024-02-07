package twosum

func TwoSum(nums []int, target int) []int {
	result := []int{}
	sum := 0
	for index, value := range nums {
		numIndices := len(result)
		increment := sum + value
		if numIndices == 2 {
			if target < increment {
				if withoutFirst := (increment - nums[result[0]]); withoutFirst == target {
					sum = withoutFirst
					result = []int{result[1], index}
				} else if withoutSecond := (increment - nums[result[1]]); withoutSecond == target {
					sum = withoutSecond
					result = []int{result[0], index}
				}
			}
		} else {
			if 0 <= target && increment <= target {
				sum = increment
				result = append(result, index)
			}
			if target < 0 && target <= increment {
				sum = increment
				result = append(result, index)
			}
		}
		if len(result) == 2 && sum == target {
			return result
		}
	}
	return []int{}
}
