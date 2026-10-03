func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m + n != len(s3) {return false}

	if m < n {
		m, n = n, m
		s1, s2 = s2, s1
	}

	dp := make([]bool, n+1)
	dp[n] = true

	for i := m; i > -1; i-- {
		nextDP := make([]bool, n + 1)
		if i == m {nextDP[n] = true}
		for j := n; j > -1; j-- {
			if i < m && s1[i] == s3[i+j] && dp[j] {
				nextDP[j] = true
			}
			if j < n && s2[j] == s3[i+j] && nextDP[j+1] {
				nextDP[j] = true
			}
		}
		dp = nextDP
	}

	return dp[0]	
}
