func longestIncreasingPath(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
    dirs := [][2]int{{0,1},{1,0},{0,-1},{-1,0}}
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
		for j := range dp[i] { dp[i][j] = -1 }
	}

	var dfs func(r, c, prevVal int) int
	dfs = func(r, c, prevVal int) int {
		if r >= rows || c >= cols || r < 0 || c < 0 ||
		prevVal >= matrix[r][c] {return 0}
		if dp[r][c] != -1 {return dp[r][c]}
		res := 1
		for _, dir := range dirs {
			res = max(res, 1 + dfs(r + dir[0], c + dir[1], matrix[r][c]))
		}
		dp[r][c] = res
		return res
	}

	maxRes := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			maxRes = max(maxRes, dfs(r, c, -1<<31))
		}
	}

	return maxRes
}
