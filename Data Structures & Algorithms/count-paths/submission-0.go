func uniquePaths(m int, n int) int {
    dirs := [][]int{{0,1}, {1,0}}

	dp := make([][]int, m + 1)
	for i := range dp { 
		dp[i] = make([]int, n + 1)
		for j := range dp[i] {dp[i][j] = -1}
	}
	
	var dfs func(int, int) int
	dfs = func(r, c int) int {
		if dp[r][c] != -1 {return dp[r][c]}
		if r == m || c == n {return 0}
		if r == m - 1 && c == n - 1 { return 1 }
		dp[r][c] = 0
		for _, dir := range dirs {
			dp[r][c] += dfs(r + dir[0], c + dir[1])
		}
		return dp[r][c]
	}

	return dfs(0, 0)
}
