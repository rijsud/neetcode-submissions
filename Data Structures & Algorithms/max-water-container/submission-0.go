func maxArea(heights []int) int {
	maxArea := 0
	maxHeight1 := [2]int{0, 0} // value, index
	maxHeight2 := [2]int{0, len(heights)-1}
	l, r := 0, len(heights) - 1
	for l < r {
		curArea := min(heights[l], heights[r]) * (r-l)
		if max(heights[l], maxHeight1[0]) != maxHeight1[0] && curArea > maxArea {
			maxHeight1[0] = heights[l]
			maxHeight1[1] = l
		}
		if max(heights[r], maxHeight2[0]) != maxHeight2[0] && curArea > maxArea {
			maxHeight2[0] = heights[r]
			maxHeight2[1] = r
		}
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
		// res := min(maxHeight1[0], maxHeight2[0]) * (maxHeight2[1]-maxHeight1[1])
		maxArea = max(curArea, maxArea)
	}
	return maxArea
}
