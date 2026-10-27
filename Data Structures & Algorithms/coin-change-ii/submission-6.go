func change(amount int, coins []int) int {
    n := len(coins)
	dp := make([]int, amount + 1)
	dp[0] = 1

	for i := n - 1; i > -1; i-- {
		for amt := 1; amt <= amount; amt++ {
			if amt - coins[i] >= 0 {
				dp[amt] += dp[amt-coins[i]]
			}
		}
	}

	return dp[amount]
}
