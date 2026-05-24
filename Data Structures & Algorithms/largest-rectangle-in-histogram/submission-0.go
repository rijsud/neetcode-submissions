func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := [][2]int{}

	for index, height := range heights {
		start := index
		for len(stack) != 0 && stack[len(stack)-1][1] > height {
			curIndex, curHeight := stack[len(stack)-1][0], stack[len(stack)-1][1]
			maxArea = max(maxArea, (index - curIndex) * curHeight)
			start = curIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{start, height})
	}

	for _, pair := range stack {
		maxArea = max(maxArea, pair[1] * (len(heights) - pair[0]))
	}

	return maxArea
}
