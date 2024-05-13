package containerwithmostwater

func maxArea(height []int) int {
	return maxAreaTwoPointers(height)
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

func maxAreaTwoPointers(height []int) int {
	if 0 == len(height) {
		return 0
	}
	var mostWaterAmount int
	leftLinePointer, rightLinePointer := 0, len(height)-1
	for leftLinePointer != rightLinePointer {
		distance, waterAmount := rightLinePointer-leftLinePointer, 0
		if isLeftLower := height[leftLinePointer] < height[rightLinePointer]; isLeftLower {
			waterAmount = height[leftLinePointer] * distance
			leftLinePointer += 1
		} else {
			waterAmount = height[rightLinePointer] * distance
			rightLinePointer -= 1
		}
		if mostWaterAmount < waterAmount {
			mostWaterAmount = waterAmount
		}
	}
	return mostWaterAmount
}
