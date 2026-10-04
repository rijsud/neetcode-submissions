func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := range dp {dp[i] = amount + 1}
	dp[0] = 0

	for amt := 1; amt <= amount; amt++ {
		for _, coin := range coins {
			if amt - coin >= 0 {
				dp[amt] = min(dp[amt], 1 + dp[amt - coin])
			}
		}
	}

	if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}
