func maxAreaOfIsland(grid [][]int) int {
    maxArea, curArea := 0, 0
	rows, cols := len(grid), len(grid[0])

	var backtrack func(int, int)
	backtrack = func(r, c int) {
		if r == rows || c == cols || r < 0 || c < 0 ||
		grid[r][c] == 0 {return}

		grid[r][c] = 0
		curArea++
		backtrack(r + 1, c)
		backtrack(r - 1, c)
		backtrack(r, c + 1)
		backtrack(r, c - 1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				curArea = 0
				backtrack(r, c)
				maxArea = max(maxArea, curArea)
			}
		}
	}

	return maxArea
}
