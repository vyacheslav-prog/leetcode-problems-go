package containerwithmostwater

func maxArea(height []int) int {
	return maxAreaBruteForce(height)
}

func maxAreaBruteForce(height []int) int {
	var mostWaterAmount int
	for leftLineIndex, leftLine := range height {
		for rightLineIndex := leftLineIndex + 1; rightLineIndex != len(height); rightLineIndex += 1 {
			rightLine, distance := height[rightLineIndex], rightLineIndex-leftLineIndex
			waterAmount := min(leftLine, rightLine) * distance
			if mostWaterAmount < waterAmount {
				mostWaterAmount = waterAmount
			}
		}
	}
	return mostWaterAmount
}
