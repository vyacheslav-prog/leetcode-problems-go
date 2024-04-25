package containerwithmostwater

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	leftLineIndex, rightLineIndex := 0, len(height)-1
	for leftLineIndex+1 < rightLineIndex && height[leftLineIndex] < height[leftLineIndex+1] {
		leftLineIndex += 1
	}
	return min(height[leftLineIndex], height[rightLineIndex]) * (rightLineIndex - leftLineIndex)
}
