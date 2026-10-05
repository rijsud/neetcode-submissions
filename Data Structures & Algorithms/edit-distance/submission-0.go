func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1))
	for i := range dp {
		dp[i] = make([]int, len(word2))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

    var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == len(word1) {return len(word2)-j}
		if j == len(word2) {return len(word1)-i}
		if dp[i][j] != -1 {return dp[i][j]}
		res := 0
		if word1[i] == word2[j] {
			res = dfs(i+1, j+1)
		} else {
			res = min(1 + dfs(i, j + 1), 
					1 + dfs(i+1, j+1),
					1 + dfs(i+1, j))
		}
		dp[i][j] = res
		return res
	}

	return dfs(0, 0)
}
