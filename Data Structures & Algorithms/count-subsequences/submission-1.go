func numDistinct(s string, t string) int {
    m, n := len(s), len(t)
	dp := make([]int, n + 1)
	dp[n] = 1

	for i := m - 1; i > -1; i-- {
		prev := 1
		for j := n - 1; j > -1; j-- {
			res := dp[j]
			if s[i] == t[j] {
				res += prev
			}
			prev = dp[j]
			dp[j] = res
		}
	}

	return dp[0]
}