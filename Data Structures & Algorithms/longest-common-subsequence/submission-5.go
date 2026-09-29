func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
	for i := range dp {dp[i] = make([]int, n+1)}

	for i := m - 1; i > -1; i-- {
		for j := n - 1; j > -1; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}
