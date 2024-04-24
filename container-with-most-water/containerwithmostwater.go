package containerwithmostwater

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	if height[1] < height[0] {
		return height[1]
	}
	return height[0]
}
