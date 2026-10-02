func change(amount int, coins []int) int {
    dp := make([]int, amount + 1)
	dp[0] = 1
	for i := len(coins) - 1; i > -1; i-- {
		for a := 1; a <= amount; a++ {
			if a - coins[i] >= 0 {
				dp[a] += dp[a-coins[i]]
			}
		}
	}

	return dp[amount]
}
