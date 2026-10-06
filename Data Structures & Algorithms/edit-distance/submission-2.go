func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m + 1)
	for i := range dp {
		dp[i] = make([]int, n + 1)
		for j := range dp[i] {dp[i][j] = -1}
	}

    var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= m {return n - j}
		if j >= n {return m - i}
		if dp[i][j] != -1 {return dp[i][j]}
		res := 0
		if word1[i] == word2[j] {
			res = dfs(i + 1, j + 1) 
		} else {
			res = 1 + min(dfs(i+1, j), dfs(i, j+1), dfs(i+1, j+1))
		}
		dp[i][j] = res
		return res
	}

	return dfs(0, 0)
}
