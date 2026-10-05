func numDistinct(s string, t string) int {
    m, n := len(s), len(t)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {dp[i][j] = -1}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j == n {return 1}
		if i == m {return 0}
		if dp[i][j] != -1 {return dp[i][j]}
		res := dfs(i + 1, j)
		if s[i] == t[j] {
			res += dfs(i + 1, j + 1)
		}
		dp[i][j] = res
		return res
	}

	return dfs(0, 0)
}
