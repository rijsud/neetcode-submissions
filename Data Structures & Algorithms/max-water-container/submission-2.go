func maxArea(heights []int) int {
	maxArea := 0
	l, r := 0, len(heights) - 1
	for l < r {
		curArea := min(heights[l], heights[r]) * (r-l)
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
		maxArea = max(curArea, maxArea)
	}
	return maxArea
}
