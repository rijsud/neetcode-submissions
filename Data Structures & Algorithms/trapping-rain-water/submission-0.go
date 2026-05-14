// func trap(height []int) int {
// 	maxArea := 0
// 	leftIndex, rightIndex := 0, len(height) - 1
// 	for leftIndex < rightIndex {
// 		curArea := min(height[leftIndex], height[rightIndex]) * (rightIndex-leftIndex)
// 		if height[leftIndex] < height[rightIndex] {
// 			leftIndex++
// 		} else {
// 			rightIndex--
// 		}
// 		maxArea = max(maxArea,curArea)
// 	}
// 	return maxArea
// }

func trap(height []int) int {
	maxArea := 0
	prefix, suffix := make([]int, len(height)), make([]int, len(height))

	maxHeight := 0
	for i, num := range height {
		if num > maxHeight {maxHeight = num}
		prefix[i] = maxHeight
	}

	maxHeight = 0
	for i := len(height) - 1; i > -1; i-- {
		if height[i] > maxHeight {maxHeight = height[i]}
		suffix[i] = maxHeight
	}

	for i, _ := range height {
		maxArea += min(prefix[i], suffix[i]) - height[i]
	}

	return maxArea
}