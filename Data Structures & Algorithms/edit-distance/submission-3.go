func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m + 1)
	for i := range dp {
		dp[i] = make([]int, n + 1)
		dp[i][n] = m - i
	}

    for j := 0; j <= n; j++ {
		dp[m][j] = n - j
	}

	for i := m - 1; i > -1; i--{
		for j := n - 1; j > -1; j-- {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = 1 + min(dp[i][j+1], dp[i+1][j], dp[i+1][j+1])
			}
		}
	}

	return dp[0][0]
}
