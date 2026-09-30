func change(amount int, coins []int) int {
	n := len(coins)
    dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, amount + 1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, amt int) int {
		if amt == 0 {return 1}
		if i >= n || amt < 0 {return 0}
		if dp[i][amt] != -1 {return dp[i][amt]}

		dp[i][amt] = dfs(i + 1, amt) + dfs(i, amt - coins[i])
		return dp[i][amt]
	}

	return dfs(0, amount)
}
