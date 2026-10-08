func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	dp := make([][]bool, m + 1)
	for i := range dp { dp[i] = make([]bool, n + 1)}
	dp[m][n] = true

	for i := m; i > -1; i-- {
		for j := n - 1; j > -1; j-- {
			match := i < m && (s[i] == p[j] || p[j] == '.')
			if j + 1 < n && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (match && dp[i+1][j])
			} else if match {
				dp[i][j] = dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}
