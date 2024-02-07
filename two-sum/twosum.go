package twosum

func TwoSum(nums []int, target int) []int {
	diffMap := map[int]int{}
	for index, value := range nums {
		if diffIndex, ok := diffMap[target-value]; ok {
			return []int{diffIndex, index}
		}
		diffMap[value] = index
	}
	return []int{}
}
