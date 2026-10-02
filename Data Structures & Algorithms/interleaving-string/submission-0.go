func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m + n != len(s3) {return false}

    dp := make([][]bool, m + 1)
	for i := range dp {dp[i] = make([]bool, n + 1)}
	dp[m][n] = true

	for i := m; i > -1; i-- {
		for j := n; j > -1; j-- {
			if i < m && s1[i] == s3[i+j] && dp[i+1][j] {
				dp[i][j] = true
			}
			if j < n && s2[j] == s3[i+j] && dp[i][j+1] {
				dp[i][j] = true
			}
		}
	}

	return dp[0][0]
}
