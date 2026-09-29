func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m < n {
		m, n = n, m
		text1, text2 = text2, text1
	}
    dp := make([]int, n+1)

	for i := m - 1; i > -1; i-- {
		prev := 0
		for j := n - 1; j > -1; j-- {
			temp := dp[j]
			if text1[i] == text2[j] {
				dp[j] = 1 + prev
			} else {
				dp[j] = max(dp[j], dp[j+1])
			}
			prev = temp
		}
	}

	return dp[0]
}
