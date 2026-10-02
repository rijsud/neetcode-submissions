func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1)+len(s2) != len(s3) {
        return false
    }

    m, n := len(s1), len(s2)
    dp := map[[2]int]bool{}

	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i == m && j == n {return true}

		key := [2]int{i,j}
		if val, ok := dp[key]; ok { return val }

		if i < m && s1[i] == s3[i+j] && dfs(i + 1, j) {
			return true
		}
		if j < n && s2[j] == s3[i+j] && dfs(i, j + 1) {
			return true
		}

		dp[key] = false
		return dp[key]
	}

	return dfs(0, 0)
}