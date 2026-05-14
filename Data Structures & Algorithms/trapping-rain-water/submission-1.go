func trap(height []int) int {
	maxArea := 0
	leftIndex, rightIndex := 0, len(height) - 1
	leftMax, rightMax := height[leftIndex], height[rightIndex]
	
	for leftIndex < rightIndex {
		curLeft, curRight := height[leftIndex], height[rightIndex]

		if curLeft < curRight {
			if leftMax < curLeft {leftMax = curLeft}
			maxArea += leftMax - curLeft
			leftIndex++
		} else {
			if rightMax < curRight {rightMax = curRight}
			maxArea += rightMax - curRight
			rightIndex--
		}
	}

	return maxArea
}
