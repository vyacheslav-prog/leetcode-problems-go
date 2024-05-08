package containerwithmostwater

func maxArea(height []int) int {
	return maxAreaBruteForce(height)
}

func maxAreaBruteForce(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var mostWaterAmount int
	for leftLineIndex := 0; leftLineIndex != len(height); leftLineIndex += 1 {
		for rightLineIndex := leftLineIndex + 1; rightLineIndex != len(height); rightLineIndex += 1 {
			waterAmount := min(height[leftLineIndex], height[rightLineIndex]) * (rightLineIndex - leftLineIndex)
			if mostWaterAmount < waterAmount {
				mostWaterAmount = waterAmount
			}
		}
	}
	return mostWaterAmount
}
