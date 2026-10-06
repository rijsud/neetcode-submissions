func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
    dp := make([][]int, m + 1)
	for i := range dp {
		dp[i] = make([]int, n + 1)
		dp[i][n] = 1
	}

	// for j := 0; j <= n; j++ {
	// 	dp[m][j] = 1
	// }

	for i := m - 1; i > -1; i-- {
		for j := n - 1; j > -1; j-- {
			dp[i][j] = dp[i+1][j]
			if s[i] == t[j] {
				dp[i][j] += dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]

}
